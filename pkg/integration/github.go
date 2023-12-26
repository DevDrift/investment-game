package integration

import (
	"io"
	"net/http"
)

type JsonRequest struct {
	Url string `json:"url"`
}

func (inputRequest *JsonRequest) GetRawData() (result []byte, err error) {
	client := http.Client{}
	resp, err := client.Get(inputRequest.Url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	result, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
