package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/MikeMwita/fedha.git/services/app-payment/models"
	"time"
)

func CreateSTKPushRequestBody(consumerKey string, consumerSecret string, baseURL string, shortcode string, passkey string, phoneNumber string, callbackURL string, accountReference string, transactionDesc string) models.STKPushRequestBody {
	timestamp := time.Now().Format("20060102150405")

	// base64 encoding of the shortcode + passkey + timestamp
	passwordToEncode := fmt.Sprintf("%s%s%s", shortcode, passkey, timestamp)
	password := base64.StdEncoding.EncodeToString([]byte(passwordToEncode))

	return models.STKPushRequestBody{
		BusinessShortCode: shortcode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            "1",
		PartyA:            phoneNumber,
		PartyB:            shortcode,
		PhoneNumber:       phoneNumber,
		CallbackURL:       callbackURL,
		AccountReference:  accountReference,
		TransactionDesc:   transactionDesc,
	}
}
