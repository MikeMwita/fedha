package config

import (
	"fmt"
	"net/http"
	"os"
)

type DefaultMpesaConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	baseUrl        string
	client         *http.Client
}

type MpesaConfig struct {
	ConsumerKey      string `json:"consumer_key"`
	ConsumerSecret   string `json:"consumer_secret"`
	BaseUrl          string `json:"base_url"`
	PhoneNumber      string `json:"phone_number"`
	Shortcode        string `json:"shortcode"`
	Passkey          string `json:"passkey"`
	CallbackURL      string `json:"callback_url"`
	AccountReference string `json:"account_reference"`
	TransactionDesc  string `json:"transaction_desc"`
}

type Config struct {
	Mpesa MpesaConfig `json:"mpesa"`
}

func LoadConfig() (*Config, error) {
	mpesaConsumerKey, ok := os.LookupEnv("MPESA_CONSUMER_KEY")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_CONSUMER_KEY")
	}

	mpesaConsumerSecret, ok := os.LookupEnv("MPESA_CONSUMER_SECRET")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_CONSUMER_SECRET")
	}

	mpesaBaseUrl, ok := os.LookupEnv("MPESA_BASE_URL")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_BASE_URL")

	}
	mpesaPhoneNumber, ok := os.LookupEnv("MPESA_PHONE_NUMBER")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_PHONE_NUMBER")
	}

	mpesaShortcode, ok := os.LookupEnv("MPESA_SHORTCODE")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_SHORTCODE")
	}

	mpesaPasskey, ok := os.LookupEnv("MPESA_PASSKEY")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_PASSKEY")
	}

	mpesaCallbackURL, ok := os.LookupEnv("MPESA_CALLBACK_URL")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_CALLBACK_URL")
	}

	mpesaAccountReference, ok := os.LookupEnv("MPESA_ACCOUNT_REFERENCE")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_ACCOUNT_REFERENCE")
	}
	mpesaTransactionDesc, ok := os.LookupEnv("MPESA_TRANSACTION_DESC")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable MPESA_TRANSACTION_DESC")
	}

	cfg := &Config{
		Mpesa: MpesaConfig{
			ConsumerKey:      mpesaConsumerKey,
			ConsumerSecret:   mpesaConsumerSecret,
			BaseUrl:          mpesaBaseUrl,
			PhoneNumber:      mpesaPhoneNumber,
			Shortcode:        mpesaShortcode,
			Passkey:          mpesaPasskey,
			CallbackURL:      mpesaCallbackURL,
			AccountReference: mpesaAccountReference,
			TransactionDesc:  mpesaTransactionDesc,
		},
	}

	return cfg, nil
}
