package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"bet-server/app"
	push "bet-server/push"
	"bet-server/repository"
)

func RegisterDevice(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)

	var body push.TokenRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Invalid body", err)
		http.Error(w, "Invalid body", 400)
		return
	}

	_, err = app.FirestoreClient.Collection("tokens").Doc(token).Set(context.Background(), body)

	if err != nil {
		log.Println("Unable to save fcm token.", err)
		http.Error(w, "Unable to save fcm token.", http.StatusBadRequest)
		return
	}
}

func UnregisterDevice(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)

	_, err := app.FirestoreClient.Collection("tokens").Doc(token).Delete(context.Background())

	if err != nil {
		log.Println("Unable to delete fcm token.", err)
		http.Error(w, "Unable to delete fcm token.", http.StatusBadRequest)
		return
	}
}

func SendNotification(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)

	var body repository.NotificationRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Invalid body", err)
		http.Error(w, "Invalid body", 400)
		return
	}

	go push.SendNotificationToUser(token, body.Title, body.Body, body.Type, body.Deeplink)
}
