package models

import (
	"testing"
)

func TestAssetRequest_Random(t *testing.T) {
	req := AssetRequest{}
	random, err := req.Random()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(random)
}
