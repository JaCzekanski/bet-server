package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_firestore "cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

const (
	ADDR = ":8080"
)

type ReturnId struct {
	ID string `json:"id"`
}

const (
	STATE_OPEN   = "OPEN"
	STATE_ACTIVE = "ACTIVE"
	STATE_CLOSED = "CLOSED"
)

type BetEntry struct {
	Bid   *int      `firestore:"bid",json:"bid"`
	Date  time.Time `firestore:"date"`
	Score string    `firestore:"score",json:"score"`
}

type Bet struct {
	State   string              `firestore:"state"`
	MatchID string              `firestore:"match_id"`
	Users   map[string]bool     `firestore:"users"`
	Bets    map[string]BetEntry `firestore:"bets"`
}

var firestore *_firestore.Client

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func getAuth(r *http.Request) *string {
	token := r.Header.Get("Authorization")
	if len(token) < 1 {
		return nil
	}
	return &token
}

func firebaseAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.Header().Set("Content-Type", "application/json")
		if getAuth(r) == nil {
			http.Error(w, "Not authorized", 401)
			return
		}
		h.ServeHTTP(w, r)
	}
}

func CreateBet(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)
	params := mux.Vars(r)

	var body BetEntry
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid body", 400)
		return
	}

	bet := Bet{
		State:   STATE_OPEN,
		MatchID: params["matchId"],
		Users: map[string]bool{
			token: true,
		},
		Bets: map[string]BetEntry{
			token: BetEntry{
				Bid:   body.Bid,
				Date:  time.Now(),
				Score: body.Score,
			},
		},
	}

	ref, _, err := firestore.Collection("bets").Add(context.Background(), bet)

	if err != nil {
		http.Error(w, "Unable to create Match record.", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(ReturnId{ID: ref.ID})
}

func PutBet(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)
	params := mux.Vars(r)

	var body BetEntry
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid body", 400)
		return
	}

	var bet Bet
	ref, err := firestore.Collection("bets").Doc(params["betId"]).Get(context.Background())

	if err != nil {
		http.Error(w, "Unable to find bet.", http.StatusBadRequest)
		return
	}

	err = ref.DataTo(&bet)
	if err != nil {
		http.Error(w, "Unable to parse bet.", http.StatusBadRequest)
		return
	}

	bet.Users[token] = true
	bet.Bets[token] = BetEntry{
		Bid:   body.Bid,
		Date:  time.Now(),
		Score: body.Score,
	}

	_, err = firestore.Collection("bets").Doc(params["betId"]).Set(context.Background(), bet)
	if err != nil {
		http.Error(w, "Unable to save bet record.", http.StatusBadRequest)
		return
	}
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

	router := mux.NewRouter()
	router.HandleFunc("/bet/{matchId}", use(CreateBet, firebaseAuth)).Methods("POST")
	router.HandleFunc("/bet/{betId}", use(PutBet, firebaseAuth)).Methods("PUT")

	log.Printf("Running server on port %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, handlers.LoggingHandler(os.Stdout, router)))
}
