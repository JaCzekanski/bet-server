package main

import (
	"fmt"
	"golang.org/x/net/context"
)

type User struct {
	Nick string `firestore:"nick"`
}

var nicknameCache = make(map[string]string)

func loadNickname(userID string) (string, error) {
	var tokenObject TokenRequest
	ref, err := firestore.Collection("users").Doc(userID).Get(context.Background())

	if err != nil {
		return "", fmt.Errorf("unable to load nick for userID %s", userID)
	}

	err = ref.DataTo(&tokenObject)
	if err != nil {
		return "", fmt.Errorf("unable to parse nick for userID %s", userID)
	}

	return tokenObject.FcmToken, nil
}

// GetUserNickname Get username from cache or load if from Firebase
func GetUserNickname(userID string) (string, error) {
	if name, exist := nicknameCache[userID]; exist {
		return name, nil
	}

	// Get nick from Firebase
	nick, err := loadNickname(userID)
	if err != nil {
		return "", fmt.Errorf("GetUserNickname error: %v", err)
	}

	nicknameCache[userID] = nick
	return nick, nil
}
