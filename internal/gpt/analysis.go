package gpt

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const apiURL = ""
const cle = ""

type GPTRequest struct {
	Prompt string `json:"prompt"`
}

type GPTResponse struct {
	Completion string `json:"completion"`
}

func AnalyzeDiff(diff string) (string, error) {
	requestBody, err := json.Marshal(GPTRequest{Prompt: diff})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+cle)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var gptResp GPTResponse
	if err := json.NewDecoder(resp.Body).Decode(&gptResp); err != nil {
		return "", err
	}

	return gptResp.Completion, nil
}
