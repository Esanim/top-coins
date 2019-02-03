package externalcoinclient

import (
	"testing"
)

func TestList(t *testing.T) {
	api, _ := New(Config{apiKey: "none"}, BaseURL(cryptoCompareBaseURL))
	listings, err := api.ListCoins()
	ok(t, err)
	if len(listings) == 0 {
		t.FailNow()
	}
}
