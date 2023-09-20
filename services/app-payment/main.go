package main

import (
	"fmt"
	"github.com/MikeMwita/fedha.git/services/app-payment/config"
	"github.com/MikeMwita/fedha.git/services/app-payment/controllers"
	"github.com/MikeMwita/fedha.git/services/app-payment/validator"
	"log"
)

const (
	BaseURL = "https://sandbox.safaricom.co.ke"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	consumerKey := cfg.Mpesa.ConsumerKey
	consumerSecret := cfg.Mpesa.ConsumerSecret
	shortcode := cfg.Mpesa.Shortcode
	passkey := cfg.Mpesa.Passkey
	phoneNumber := cfg.Mpesa.PhoneNumber
	callbackURL := cfg.Mpesa.CallbackURL
	accountReference := cfg.Mpesa.AccountReference
	transactionDesc := cfg.Mpesa.TransactionDesc

	mpesa := controllers.NewMpesa(&controllers.MpesaOpts{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		BaseURL:        BaseURL,
	})

	accessTokenResponse, err := mpesa.GenerateAccessToken()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", accessTokenResponse)

	stkPushReq := controllers.CreateSTKPushRequestBody(consumerKey, consumerSecret, BaseURL, shortcode, passkey, phoneNumber, callbackURL, accountReference, transactionDesc)
	err = validator.ValidateSTKPushRequestBody(&stkPushReq)
	if err != nil {
		log.Fatalln(err)
	}
	response, err := mpesa.InitiateSTKPushRequest(&stkPushReq)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", response)
}
