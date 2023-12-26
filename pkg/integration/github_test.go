package integration

import (
	"encoding/json"
	"os"
	"testing"
)

func TestJsonRequest_GetJsonRawData(t *testing.T) {
	req := JsonRequest{
		Url: "https://raw.githubusercontent.com/DevDrift/investment-game/main/data/names/cryptocurrency.json",
	}
	data, err := req.GetRawData()
	if err != nil {
		t.Error(err)
		return
	}
	var list []string
	err = json.Unmarshal(data, &list)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(len(list))
}

func TestJsonRequest_GetImageRawData(t *testing.T) {
	req := JsonRequest{
		Url: "https://raw.githubusercontent.com/DevDrift/investment-game/images/cryptocurrency/1.png",
	}
	data, err := req.GetRawData()
	if err != nil {
		t.Error(err)
		return
	}
	create, err := os.Create("1.png")
	if err != nil {
		return
	}
	_, err = create.Write(data)
	if err != nil {
		t.Error(err)
		return
	}
}
