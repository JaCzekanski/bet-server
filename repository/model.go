package repository

import (
	"time"
)

const (
	StateBefore = "BEFORE"
	StateDuring = "DURING"
	StateAfter  = "AFTER"
)

type ReturnID struct {
	ID string `json:"id"`
}

type BetEntry struct {
	Bid   *int      `firestore:"bid" json:"bid"`
	Date  time.Time `firestore:"date"`
	Score string    `firestore:"score" json:"score"`
}

type Bet struct {
	betID   string              // skip serialization
	State   string              `firestore:"state"`
	MatchID string              `firestore:"matchId"`
	Users   map[string]bool     `firestore:"users"`
	Bets    map[string]BetEntry `firestore:"bets"`
}


type Match struct {
	Team1 string  `firestore:"team1"`
	Team2 string  `firestore:"team2"`
	Score *string `firestore:"score"`
	State string  `firestore:"state"`
}

type NotificationRequest struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Type     string `json:"type"`
	Deeplink string `json:"deeplink"`
}
