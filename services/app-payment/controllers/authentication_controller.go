package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Mpesa struct {
	consumerKey    string
	consumerSecret string
	baseURL        string
	client         *http.Client
}

type MpesaOpts struct {
	ConsumerKey    string
	ConsumerSecret string
	BaseURL        string
}

type MpesaAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RequestID    string `json:"requestId"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewMpesa(m *MpesaOpts) *Mpesa {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &Mpesa{
		consumerKey:    m.ConsumerKey,
		consumerSecret: m.ConsumerSecret,
		baseURL:        m.BaseURL,
		client:         client,
	}
}

func (m *Mpesa) MakeRequest(req *http.Request) ([]byte, error) {
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (m *Mpesa) GenerateAccessToken() (*MpesaAccessTokenResponse, error) {
	url := fmt.Sprintf("%s/oauth/v1/generate?grant_type=client_credentials", m.baseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(m.consumerKey, m.consumerSecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	accessTokenResponse := new(MpesaAccessTokenResponse)
	if err := json.Unmarshal(resp, &accessTokenResponse); err != nil {
		return nil, err
	}

	return accessTokenResponse, nil
}
