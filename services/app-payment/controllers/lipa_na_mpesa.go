package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MikeMwita/fedha.git/services/app-payment/models"
	"net/http"
)

func (m *Mpesa) InitiateSTKPushRequest(body *models.STKPushRequestBody) (*models.STKPushRequestResponse, error) {
	url := fmt.Sprintf("%s/mpesa/stkpush/v1/processrequest", m.baseURL)

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	accessTokenResponse, err := m.GenerateAccessToken()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessTokenResponse.AccessToken))

	resp, err := m.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	stkPushResponse := new(models.STKPushRequestResponse)
	if err := json.Unmarshal(resp, &stkPushResponse); err != nil {
		return nil, err
	}

	return stkPushResponse, nil
}
