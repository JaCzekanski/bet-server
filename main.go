package main

import (
	_firestore "cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

const (
	ADDR = ":8080"
)

var firestore *_firestore.Client
var fcmClient *messaging.Client

func createRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/bet/{matchId}", use(CreateBet, firebaseAuth)).Methods("POST")
	router.HandleFunc("/bet/{betId}", use(PutBet, firebaseAuth)).Methods("PUT")
	router.HandleFunc("/bet/{betId}", use(DeleteBet, firebaseAuth)).Methods("DELETE")
	router.HandleFunc("/bet/{betId}/invite/{userId}", use(InviteUserToBet, firebaseAuth)).Methods("POST")
	router.HandleFunc("/changeUserInBet/{betId}/from/{oldId}/to/{newId}", use(ChangeUserInBet, firebaseAuth)).Methods("POST")
	router.HandleFunc("/register", use(RegisterDevice, firebaseAuth)).Methods("POST")
	router.HandleFunc("/register", use(UnregisterDevice, firebaseAuth)).Methods("DELETE")
	router.HandleFunc("/notitfication", use(SendNotification, firebaseAuth)).Methods("POST")
	return router
}

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("bet-app-bc625-firebase-adminsdk-j2r9e-c7205cb679.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	firestore, err = app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	defer firestore.Close()
	log.Printf("Connected to Firebase")

	fcmClient, err = app.Messaging(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Connected to FCM")

	log.Printf("Running server on port %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, handlers.LoggingHandler(os.Stdout, handlers.ProxyHeaders(createRouter()))))
}
