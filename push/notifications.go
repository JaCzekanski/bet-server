package push

import (
	"fmt"
	"log"
	"bet-server/repository"
	"bet-server/country"
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

type Bet struct {
	BetID string
	Bet   repository.Bet
}

type Match struct {
	MatchID string
	Match   repository.Match
}

func SendInviteNotification(fromID string, toID string, bet Bet) {
	fromNick, err := repository.GetUserNickname(fromID)
	if err != nil {
		log.Println("Unable to load nickname.", err)
		return
	}

	match, err := repository.LoadMatch(bet.Bet.MatchID)
	if err != nil {
		log.Println("Unable to load match.", err)
		return
	}

	title := fmt.Sprintf("Zaproszenie od %s", fromNick)
	body := fmt.Sprintf("Obstaw wynik meczu %s:%s", country.MapCodeToCountry(match.Team1), country.MapCodeToCountry(match.Team2))
	deeplink := fmt.Sprintf(UriBet, bet.BetID)

	SendNotificationToUser(toID, title, body, TypeInvite, deeplink)
}

func SendMatchScoreInfo(match Match) {
	// 1. Get bets with given match_id
	// 2. Compose message
	// 3. Get all tokens
	// 4. Create mass push
	bets, err := LoadBets(match.MatchID)

	if err != nil {
		log.Panicln(err)
		return
	}

	for betId, bet := range *bets {
		for userId, betEntry := range bet.Bets {
			if betEntry.Bid != nil {
				var body string

				if betEntry.Score == *match.Match.Score {
					body = "Trafiłeś zakład"
				} else {
					body = "Nie trafiłeś zakładu"
				}

				title := fmt.Sprintf("Wynik meczu %s:%s    %s", country.MapCodeToCountry(match.Match.Team1), country.MapCodeToCountry(match.Match.Team2), *match.Match.Score)
				deeplink := fmt.Sprintf(UriBet, betId)

				SendNotificationToUser(userId, title, body, TypeScore, deeplink)
			}
		}
	}
}
