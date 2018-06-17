package repository

import (
	"fmt"
	"golang.org/x/net/context"
	"bet-server/app"
)

func LoadMatch(matchId string) (*Match, error) {
	var match Match
	ref, err := app.FirestoreClient.Collection("matches").Doc(matchId).Get(context.Background())

	if err != nil {
		return nil, fmt.Errorf("unable to find match %v", err)
	}

	err = ref.DataTo(&match)
	if err != nil {
		return nil, fmt.Errorf("unable to parse match %v", err)
	}

	return &match, nil
}
