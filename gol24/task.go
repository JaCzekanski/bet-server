package gol24

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"4d63.com/tz"
	"bet-server/country"
	"bet-server/app"
	"bet-server/push"
	"bet-server/repository"
)

const (
	// MS 2018
	APIURL = "http://www.gol24.pl/ajax/serwis_specjalny/serwisy_sportowe/wyniki_meczow/?strona=1&region=wszystkie&liczba=1000&data_od=2018-6-14+0%3A0%3A0&data_do=2018-7-15+23%3A59%3A59&order=phase-day&druzyna=&id_dyscypliny=&id_lig=280&na_zywo="
	// APIURL = "http://www.gol24.pl/ajax/serwis_specjalny/serwisy_sportowe/wyniki_meczow/?id=1&typ=dyscyplina&na_zywo=1&strona=1&liczba=7&region=&id_lig%5B%5D=5&id_lig%5B%5D=1&id_lig%5B%5D=3&id_lig%5B%5D=59&id_lig%5B%5D=63&id_lig%5B%5D=65&id_lig%5B%5D=121&id_lig%5B%5D=41&id_lig%5B%5D=49&id_lig%5B%5D=43&id_lig%5B%5D=55&id_lig%5B%5D=51&id_lig%5B%5D=53"
)

const (
	StateApiBefore = "przd"
	StateApiAfter  = "po"
)

const (
	StateBefore = "BEFORE"
	StateDuring = "DURING"
	StateAfter  = "AFTER"
)

func mapState(apiState string) string {
	switch apiState {
	case StateApiBefore:
		return StateBefore
	case StateApiAfter:
		return StateAfter
	default: // dunno what state it is (during)
		return StateDuring
	}
}

func downloadMatches() response {
	log.Println("Downloading matches...")
	resp, err := http.Get(APIURL)
	if err != nil {
		panic(err)
	}

	var data response

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(contents, &data); err != nil {
		panic(err)
	}

	return data
}

func uploadLeaguesData(data response) {
	log.Println("Uploading leagues data...")
	batch := app.FirestoreClient.Batch()

	for _, element := range data.Refs.Leagues {
		log.Printf("Uploading document %d\n", element.ID)
		batch.Set(app.FirestoreClient.Collection("leagues").Doc(strconv.Itoa(element.ID)), map[string]interface{}{
			"name": element.Name,
		})
	}

	_, err := batch.Commit(context.Background())
	if err != nil {
		panic(err)
	}
}

func prepareMatchesData(data response) map[string]repository.Match {
	location, err := tz.LoadLocation("Poland")
	if err != nil {
		panic(err)
	}

	matches := make(map[string]repository.Match)

	for _, element := range data.Data {
		date, _ := time.ParseInLocation("2006-01-02 15:04:05", element.Date, location)
		team1 := data.Refs.Teams[strconv.Itoa(element.Scores[0].Team)]
		team2 := data.Refs.Teams[strconv.Itoa(element.Scores[1].Team)]

		var score *string

		if element.Scores[0].Score != nil && element.Scores[1].Score != nil {
			s := fmt.Sprintf("%s:%s", *element.Scores[0].Score, *element.Scores[1].Score)
			score = &s
		}

		matches[strconv.Itoa(element.ID)] = repository.Match{
			Event: element.LeagueID,
			Score: score,
			Team1: strings.ToLower(country.MapCountryToIso(team1.Name)),
			Team2: strings.ToLower(country.MapCountryToIso(team2.Name)),
			Date:  date,
			State: mapState(element.EventStatus),
		}
	}

	return matches
}

func calculateDiff(old map[string]repository.Match, new map[string]repository.Match) map[string]repository.Match {
	if old == nil {
		log.Printf("No previous results\n")
		return new
	}

	diffed := make(map[string]repository.Match)

	for id, newElement := range new {
		oldElement, oldExist := old[id]

		if !oldExist {
			log.Printf("Document %s new\n", id)
		} else if !reflect.DeepEqual(newElement, oldElement) {
			log.Printf("Document %s changed\n", id)

			// On state DURING -> AFTER  - send push to all users in bets with THIS bet_id
			log.Printf("Old: %+v\n", oldElement)
			log.Printf("New: %+v\n", newElement)

			if newElement.State == StateAfter && oldElement.State == StateDuring {
				log.Println("Match ended, sending pushes!")
				go push.SendMatchScoreInfo(push.Match{
					MatchID: id,
					Match: newElement,
				})
			}

		} else {
			continue
		}

		diffed[id] = newElement
	}
	return diffed
}

func uploadData(collection string, matches map[string]repository.Match) {
	log.Printf("Uploading %s data...\n", collection)
	batch := app.FirestoreClient.Batch()

	for id, element := range matches {
		log.Printf("Uploading document %s\n", id)
		batch.Set(app.FirestoreClient.Collection(collection).Doc(id), element)
	}

	_, err := batch.Commit(context.Background())
	if err != nil {
		panic(err)
	}
}

var previousMatches map[string]repository.Match

func DownloadDataAndUploadToFirebase() {
	// log.Println("Running downloadDataAndUploadToFirebase")
	data := downloadMatches()

	// uploadLeaguesData(data)

	matches := prepareMatchesData(data)
	diffed := calculateDiff(previousMatches, matches)

	if len(diffed) == 0 {
		log.Println("No changes in match data. Skipping upload.")
	} else {
		uploadData("matches", diffed)
	}

	previousMatches = matches
}
