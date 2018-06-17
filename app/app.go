package app

import (
	_firestore "cloud.google.com/go/firestore"
	"firebase.google.com/go/messaging"
)

var FirestoreClient *_firestore.Client
var FcmClient *messaging.Client
