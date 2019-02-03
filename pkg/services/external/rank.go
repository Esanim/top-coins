package externalcoinclient

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	api "github.com/esanim/top-coins/api/external"
)

const (
	cryptoCompareBaseURL = "https://min-api.cryptocompare.com"
)

//ExternalRankingsService interface
type ExternalRankingsService interface {
	List(context.Context) ([]api.CoinRanking, error)
}

func getCoins(ds *api.CoinsRoot) ([]api.CoinRanking, error) {
	var coins []api.CoinRanking
	for _, value := range ds.Data {
		coins = append(coins, value)
	}

	return coins, nil
}

//ListCoins retrieves coins data and their rankings (sortOrder) from cryptocompare.
func (s *Client) ListCoins() ([]api.CoinRanking, error) {
	var params []string
	url := fmt.Sprintf("%s/data/all/coinlist?%s", cryptoCompareBaseURL, strings.Join(params, "&"))

	body, err := s.makeReq(url, nil)
	resp := new(api.CoinsRoot)
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	coins, err := getCoins(resp)
	if err != nil {
		return nil, err
	}

	return coins, err
}

// NewCryptoCompareClient returns a new client instance for Crypto Compare Api
func NewCryptoCompareClient() (*Client, error) {
	return New(Config{"none"}, BaseURL(cryptoCompareBaseURL))
}
