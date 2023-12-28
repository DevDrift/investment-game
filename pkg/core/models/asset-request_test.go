package models

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssetRequest(t *testing.T) {
	req := AssetRequest{}
	random, err := req.Random()
	if err != nil {
		t.Error(err)
		return
	}
	random.Id = testId
	req.Asset = random
	addAsset, err := req.Add()
	if err != nil {
		t.Error(err)
		return
	}
	getAsset, err := req.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, addAsset.Name, getAsset.Name)
	list, err := req.List()
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, 0, len(list))
	getAsset.Id = uuid.NewString()
	req.Asset = getAsset
	updateAsset, err := req.Update([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	updateAsset, err = req.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, updateAsset.Id, testId)
	err = req.Delete([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	deletedAsset, err := req.Get([]byte(testId))
	if err != nil {
		t.Error(err)
		return
	}
	assert.NotEqual(t, deletedAsset, 0)
}
