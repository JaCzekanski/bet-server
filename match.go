package main

import (
	"fmt"
	"golang.org/x/net/context"
)

type Match struct {
	Team1 string  `firestore:"team1"`
	Team2 string  `firestore:"team2"`
	Score *string `firestore:"score"`
	State string  `firestore:"state"`
}

func LoadMatch(matchId string) (*Match, error) {
	var match Match
	ref, err := firestore.Collection("matches").Doc(matchId).Get(context.Background())

	if err != nil {
		return nil, fmt.Errorf("unable to find match %v", err)
	}

	err = ref.DataTo(&match)
	if err != nil {
		return nil, fmt.Errorf("unable to parse match %v", err)
	}

	return &match, nil
}
