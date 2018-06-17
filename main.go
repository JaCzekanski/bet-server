package main

import (
	"context"
	"firebase.google.com/go"
	"github.com/gorilla/handlers"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"bet-server/app"
	"bet-server/api"
	"bet-server/gol24"
	"github.com/jasonlvhit/gocron"
)

const (
	ADDR = ":8080"
)

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("bet-app-bc625-firebase-adminsdk-j2r9e-c7205cb679.json")
	firebaseApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	app.FirestoreClient, err = firebaseApp.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	defer app.FirestoreClient.Close()
	log.Printf("Connected to Firebase")

	app.FcmClient, err = firebaseApp.Messaging(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Connected to FCM")

	log.Printf("Running Gol24 client in background")
	gocron.Every(1).Minute().Do(gol24.DownloadDataAndUploadToFirebase)
	gocron.Start()

	log.Printf("Running server on port %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, handlers.LoggingHandler(os.Stdout, handlers.ProxyHeaders(api.CreateRouter()))))
}
