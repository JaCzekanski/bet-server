package main

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
	BetID   string
	State   string              `firestore:"state"`
	MatchID string              `firestore:"matchId"`
	Users   map[string]bool     `firestore:"users"`
	Bets    map[string]BetEntry `firestore:"bets"`
}

type TokenRequest struct {
	FcmToken string `firestore:"fcmToken" json:"fcmToken"`
}

type NotificationRequest struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Type     string `json:"type"`
	Deeplink string `json:"deeplink"`
}
