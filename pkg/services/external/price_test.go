package externalcoinclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	api "github.com/esanim/top-coins/api/external"
)

func TestGetCryptoMarketPricesFake(t *testing.T) {
	url := coinMarketCapURL + "/cryptocurrency/listings/latest?limit=1"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), url)

		quote1 := &api.Quote{Price: 1}
		quotes := make(map[string]*api.Quote)
		quotes["USD"] = quote1
		listing1 := api.PriceListing{ID: 1, Name: "Bitcoin", Symbol: "BTC", Quote: quotes}
		j, err := json.Marshal(&api.Response{Status: api.Status{}, Data: []api.PriceListing{listing1}})
		if err != nil {
			fmt.Println(err)
		}
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(string(j))),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	api, _ := New(Config{apiKey: "none"}, BaseURL(coinMarketCapURL), HTTPClient(client))
	listings, err := api.GetLatestCryptocurrencyListing(&CryptocurrencyListingsLatestOptions{
		Limit: 1,
	})
	ok(t, err)
	if len(listings) == 0 {
		t.FailNow()
	}
	if listings[0].Name != "Bitcoin" {
		t.FailNow()
	}
	if listings[0].Quote["USD"].Price <= 0 {
		t.FailNow()
	}
}

// can be tested if you have os env header apki key variable or you hardcode it here
// func TestGetCryptoMarketPricesReal(t *testing.T) {
// 	api, _ := New(Config{apiKey: "your key here"}, BaseURL(coinMarketCapURL))
// 	listings, err := api.GetLatestCryptocurrencyListing(&CryptocurrencyListingsLatestOptions{
// 		Limit: 1,
// 	})
// 	ok(t, err)
// 	if len(listings) == 0 {
// 		t.FailNow()
// 	}
// 	if listings[0].Name != "Bitcoin" {
// 		t.FailNow()
// 	}
// 	if listings[0].Quote["USD"].Price <= 0 {
// 		t.FailNow()
// 	}
// }
