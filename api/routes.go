package api

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/bet/{matchId}", Use(CreateBet, FirebaseAuth)).Methods("POST")
	router.HandleFunc("/bet/{betId}", Use(PutBet, FirebaseAuth)).Methods("PUT")
	router.HandleFunc("/bet/{betId}", Use(DeleteBet, FirebaseAuth)).Methods("DELETE")
	router.HandleFunc("/bet/{betId}/invite/{userId}", Use(InviteUserToBet, FirebaseAuth)).Methods("POST")
	router.HandleFunc("/changeUserInBet/{betId}/from/{oldId}/to/{newId}", Use(ChangeUserInBet, FirebaseAuth)).Methods("POST")
	router.HandleFunc("/register", Use(RegisterDevice, FirebaseAuth)).Methods("POST")
	router.HandleFunc("/register", Use(UnregisterDevice, FirebaseAuth)).Methods("DELETE")
	router.HandleFunc("/notification", Use(SendNotification, FirebaseAuth)).Methods("POST")
	router.HandleFunc("/notitfication", Use(SendNotification, FirebaseAuth)).Methods("POST") // Deprecated
	return router
}