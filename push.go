package main

import (
	"firebase.google.com/go/messaging"
	"fmt"
	"golang.org/x/net/context"
	"log"
)

var tokenCache = make(map[string]string)

func loadTokenForUser(userID string) (string, error) {
	var tokenObject TokenRequest
	ref, err := firestore.Collection("tokens").Doc(userID).Get(context.Background())

	if err != nil {
		return "", fmt.Errorf("unable to load token for user %s", userID)
	}

	err = ref.DataTo(&tokenObject)
	if err != nil {
		return "", fmt.Errorf("unable to parse token for user %s", userID)
	}

	return tokenObject.FcmToken, nil
}

// GetTokenForUser Get token from cache or load it from Firebase
func GetTokenForUser(userID string) (string, error) {
	if token, exist := tokenCache[userID]; exist {
		return token, nil
	}

	// Get token from Firebase
	token, err := loadTokenForUser(userID)
	if err != nil {
		return "", fmt.Errorf("sendNotificationToUser error: %v", err)
	}

	tokenCache[userID] = token
	return token, nil
}

func SendNotificationToUser(userID string, title string, body string, notificationType string, deeplink string) {
	token, err := GetTokenForUser(userID)
	if err != nil {
		log.Println("Error when getting token: ", err)
		return
	}

	push := &messaging.Message{
		Token: token,
		Data: map[string]string{
			"title":    title,
			"body":     body,
			"type":     notificationType,
			"deeplink": deeplink,
		},
	}

	_, err = fcmClient.Send(context.Background(), push)
	if err != nil {
		log.Println("Error when sending push notification: ", err)
		return
	}
}
