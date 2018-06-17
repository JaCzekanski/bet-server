package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

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
			token: {
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
	token := *getAuth(r)
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
	bet.betID = betId
	go SendInviteNotification(token, userId, bet)
}
