package models

import "net/http"

type Mpesa struct {
	consumerKey    string
	consumerSecret string
	baseURL        string
	client         *http.Client
}
