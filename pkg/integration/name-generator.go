package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GenRequest struct {
	Description string `json:"description"`
	Action      string `json:"action"`
	Type        string `json:"type"`
}

type GenResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Content string   `json:"content"`
	Data    []string `json:"data"`
}

// Generate gen new company name
func (inputReq *GenRequest) Generate() (result *GenResponse, err error) {
	url := "https://websites.automizely.com/v1/public/aigc/name-generator"
	inBody, _ := json.Marshal(inputReq)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(inBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}
	return
}
