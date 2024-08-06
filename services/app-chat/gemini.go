package app_chat

import (
	"context"
	"log"
	"net/http"
	"time"
)

func GetGeminiResponse(prompt string) (interface{}, int, error) {
	log.Println("The prompt text received :", prompt)

	prompt += ".Answer in less than 20 words"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := GenModel.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	log.Println("The response recived is : ", resp)
	var queryResponse interface{}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		queryResponse := resp.Candidates[0].Content.Parts[0]
		return queryResponse, http.StatusOK, nil
	}

	return queryResponse, http.StatusOK, nil

}
