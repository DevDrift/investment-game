package integration

import (
	"testing"
)

func TestGenRequest_Generate(t *testing.T) {
	req := GenRequest{
		Description: "Аренда коммерческой недвижимости, высотное здание торговые площади",
		Action:      "Generate",
		Type:        "store",
	}
	generate, err := req.Generate()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(generate.Data)
}
