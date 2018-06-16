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

type ReturnID struct {
	ID string `json:"id"`
}

const (
	StateBefore = "BEFORE"
	StateDuring = "DURING"
	StateAfter  = "AFTER"
)

type BetEntry struct {
	Bid   *int      `firestore:"bid" json:"bid"`
	Date  time.Time `firestore:"date"`
	Score string    `firestore:"score" json:"score"`
}

type Bet struct {
	State   string              `firestore:"state"`
	MatchID string              `firestore:"matchId"`
	Users   map[string]bool     `firestore:"users"`
	Bets    map[string]BetEntry `firestore:"bets"`
}

type TokenRequest struct {
	FcmToken string `firestore:"fcmToken" json:"fcmToken"`
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
			log.Println("Not authorized")
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
		log.Println("Invalid body", err)
		http.Error(w, "Invalid body", 400)
		return
	}

	bet := Bet{
		State:   StateBefore,
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
		log.Println("Unable to create Match record", err)
		http.Error(w, "Unable to create Match record.", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(ReturnID{ID: ref.ID})
}

func PutBet(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)
	params := mux.Vars(r)

	var body BetEntry
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Invalid body", err)
		http.Error(w, "Invalid body", 400)
		return
	}

	var bet Bet
	ref, err := firestore.Collection("bets").Doc(params["betId"]).Get(context.Background())

	if err != nil {
		log.Println("Unable to find bet.", err)
		http.Error(w, "Unable to find bet.", http.StatusBadRequest)
		return
	}

	err = ref.DataTo(&bet)
	if err != nil {
		log.Println("Unable to parse bet.", err)
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
		log.Println("Unable to save bet record.", err)
		http.Error(w, "Unable to save bet record.", http.StatusBadRequest)
		return
	}
}

func ChangeUserInBet(w http.ResponseWriter, r *http.Request) {
	// token := *getAuth(r)
	params := mux.Vars(r)

	betId := params["betId"]
	oldId := params["oldId"]
	newId := params["newId"]

	var bet Bet
	ref, err := firestore.Collection("bets").Doc(betId).Get(context.Background())

	if err != nil {
		log.Println("Unable to find bet.", err)
		http.Error(w, "Unable to find bet.", http.StatusBadRequest)
		return
	}

	err = ref.DataTo(&bet)
	if err != nil {
		log.Println("Unable to parse bet.", err)
		http.Error(w, "Unable to parse bet.", http.StatusBadRequest)
		return
	}

	bet.Users[newId] = bet.Users[oldId]
	bet.Bets[newId] = bet.Bets[oldId]

	delete(bet.Users, oldId)
	delete(bet.Bets, oldId)

	_, err = firestore.Collection("bets").Doc(params["betId"]).Set(context.Background(), bet)
	if err != nil {
		log.Println("Unable to save bet record.", err)
		http.Error(w, "Unable to save bet record.", http.StatusBadRequest)
		return
	}
}

func DeleteBet(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)
	params := mux.Vars(r)

	var bet Bet
	ref, err := firestore.Collection("bets").Doc(params["betId"]).Get(context.Background())

	if err != nil {
		log.Println("Unable to find bet.", err)
		http.Error(w, "Unable to find bet.", http.StatusBadRequest)
		return
	}

	err = ref.DataTo(&bet)
	if err != nil {
		log.Println("Unable to parse bet.", err)
		http.Error(w, "Unable to parse bet.", http.StatusBadRequest)
		return
	}

	delete(bet.Users, token)
	delete(bet.Bets, token)

	if len(bet.Users) != 0 {
		_, err = firestore.Collection("bets").Doc(params["betId"]).Set(context.Background(), bet)
		if err != nil {
			log.Println("Unable to delete bet record.", err)
			http.Error(w, "Unable to delete bet record.", http.StatusBadRequest)
			return
		}
	} else {
		_, err = firestore.Collection("bets").Doc(params["betId"]).Delete(context.Background())
		if err != nil {
			log.Println("Unable to delete bet record.", err)
			http.Error(w, "Unable to delete bet record.", http.StatusBadRequest)
			return
		}
	}
}

func InviteUserToBet(w http.ResponseWriter, r *http.Request) {
	// token := *getAuth(r)
	params := mux.Vars(r)

	betId := params["betId"]
	userId := params["userId"]

	var bet Bet
	ref, err := firestore.Collection("bets").Doc(betId).Get(context.Background())

	if err != nil {
		log.Println("Unable to find bet.", err)
		http.Error(w, "Unable to find bet.", http.StatusBadRequest)
		return
	}

	err = ref.DataTo(&bet)
	if err != nil {
		log.Println("Unable to parse bet.", err)
		http.Error(w, "Unable to parse bet.", http.StatusBadRequest)
		return
	}

	if _, exist := bet.Users[userId]; exist {
		log.Println("User already invited.")
		http.Error(w, "User already invited.", http.StatusBadRequest)
		return
	}

	bet.Users[userId] = true
	bet.Bets[userId] = BetEntry{
		Bid:   nil,
		Date:  time.Now(),
		Score: "",
	}

	_, err = firestore.Collection("bets").Doc(betId).Set(context.Background(), bet)
	if err != nil {
		log.Println("Unable to save bet record.", err)
		http.Error(w, "Unable to save bet record.", http.StatusBadRequest)
		return
	}

	// Send push notification
}

func RegisterDevice(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)

	var body TokenRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Invalid body", err)
		http.Error(w, "Invalid body", 400)
		return
	}

	_, err = firestore.Collection("tokens").Doc(token).Set(context.Background(), body)

	if err != nil {
		log.Println("Unable to save fcm token.", err)
		http.Error(w, "Unable to save fcm token.", http.StatusBadRequest)
		return
	}
}

func UnregisterDevice(w http.ResponseWriter, r *http.Request) {
	token := *getAuth(r)

	_, err := firestore.Collection("tokens").Doc(token).Delete(context.Background())

	if err != nil {
		log.Println("Unable to delete fcm token.", err)
		http.Error(w, "Unable to delete fcm token.", http.StatusBadRequest)
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
	router.HandleFunc("/bet/{betId}", use(DeleteBet, firebaseAuth)).Methods("DELETE")
	router.HandleFunc("/bet/{betId}/invite/{userId}", use(InviteUserToBet, firebaseAuth)).Methods("POST")
	router.HandleFunc("/changeUserInBet/{betId}/from/{oldId}/to/{newId}", use(ChangeUserInBet, firebaseAuth)).Methods("POST")
	router.HandleFunc("/register", use(RegisterDevice, firebaseAuth)).Methods("POST")
	router.HandleFunc("/register", use(UnregisterDevice, firebaseAuth)).Methods("DELETE")

	log.Printf("Running server on port %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, handlers.LoggingHandler(os.Stdout, handlers.ProxyHeaders(router))))
}
