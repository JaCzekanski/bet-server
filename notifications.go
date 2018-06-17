package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	TypeInvite   = "INVITE"
	TypeReminder = "REMINDER"
	TypeScore    = "SCORE"
	TypeOther    = "OTHER"
)

const (
	UriBet = "https://bet.czekanski.info/bet/%s"
)

var countries = parseCountries()

func mapCodeToCountry(code string) string {
	code = strings.ToUpper(code)
	for _, e := range countries {
		if e.Code == code {
			return e.NamePl
		}
	}
	fmt.Printf("Unable to find country %s\n", code)
	return code
}

func SendInviteNotification(fromID string, toID string, bet Bet) {
	fromNick, err := GetUserNickname(fromID)
	if err != nil {
		log.Println("Unable to load nickname.", err)
		return
	}

	match, err := LoadMatch(bet.MatchID)
	if err != nil {
		log.Println("Unable to load match.", err)
		return
	}

	title := fmt.Sprintf("Zaproszenie od %s", fromNick)
	body := fmt.Sprintf("Obstaw wynik meczu %s:%s", mapCodeToCountry(match.Team1), mapCodeToCountry(match.Team2))
	deeplink := fmt.Sprintf(UriBet, bet.betID)

	SendNotificationToUser(toID, title, body, TypeInvite, deeplink)
}
