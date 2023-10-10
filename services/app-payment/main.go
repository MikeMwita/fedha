package main

import (
	"fmt"
	"github.com/MikeMwita/fedha.git/services/app-payment/controllers"
	"github.com/MikeMwita/fedha.git/services/app-payment/validator"
	"github.com/caarlos0/env/v9"
	"log"
)

const (
	BaseURL = "https://sandbox.safaricom.co.ke"
)

type Conf struct {
	ConsumerKey      string `env:"MPESA_CONSUMER_KEY"`
	ConsumerSecret   string `env:"MPESA_CONSUMER_SECRET"`
	BaseURL          string `env:"MPESA_BASE_URL"          envDefault:"https://sandbox.safaricom.co.ke"`
	PrometheusURL    string `env:"MO_PROMETHEUS_URL"       envDefault:""`
	shortcode        string `env:"MPESA_SHORTCODE"`
	passkey          string `env:"MPESA_PASSKEY"`
	phoneNumber      string `env:"MPESA_PHONE_NUMBER"`
	callbackURL      string `env:"MPESA_CALLBACK_URL"`
	accountReference string `env:"MPESA_ACCOUNT_REFERENCE"`
	transactionDesc  string `env:"MPESA_TRANSACTION_DESC"`
}

func main() {

	cf := Conf{}
	if err := env.Parse(&cf); err != nil {
		log.Fatalf("failed to load configuration : %s", err)
	}

	//cfg, err := config.LoadConfig()
	//if err != nil {
	//	log.Fatalf("Error loading config: %v", err)
	//}
	//
	//consumerKey := cfg.Mpesa.ConsumerKey
	//consumerSecret := cfg.Mpesa.ConsumerSecret
	//shortcode := cfg.Mpesa.Shortcode
	//passkey := cfg.Mpesa.Passkey
	//phoneNumber := cfg.Mpesa.PhoneNumber
	//callbackURL := cfg.Mpesa.CallbackURL
	//accountReference := cfg.Mpesa.AccountReference
	//transactionDesc := cfg.Mpesa.TransactionDesc

	mpesa := controllers.NewMpesa(&controllers.MpesaOpts{
		ConsumerKey:    cf.ConsumerKey,
		ConsumerSecret: cf.ConsumerSecret,
		BaseURL:        BaseURL,
	})

	accessTokenResponse, err := mpesa.GenerateAccessToken()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", accessTokenResponse)

	stkPushReq := controllers.CreateSTKPushRequestBody(cf.ConsumerKey, cf.ConsumerSecret, BaseURL, cf.shortcode, cf.passkey, cf.phoneNumber, cf.callbackURL, cf.accountReference, cf.transactionDesc)
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
