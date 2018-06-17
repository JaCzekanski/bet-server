package push


type TokenRequest struct {
	FcmToken string `firestore:"fcmToken" json:"fcmToken"`
}
