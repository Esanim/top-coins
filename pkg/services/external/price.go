package externalcoinclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	api "github.com/esanim/top-coins/api/external"
)

const (
	coinMarketCapURL = "https://pro-api.coinmarketcap.com/v1"
	apiHeaderKeyName = "X-CMC_PRO_API_KEY"
)

var (
	// ErrCouldNotCast could not cast error
	ErrCouldNotCast = errors.New("could not cast")
)

// CryptocurrencyListingsLatestOptions options
type CryptocurrencyListingsLatestOptions struct {
	Limit int
}

// ExternalPriceService interface
type ExternalPriceService interface {
	CryptocurrencyListingsLatest(options *CryptocurrencyListingsLatestOptions) ([]*api.PriceListing, error)
}

// GetLatestCryptocurrencyListing gets a paginated list of all cryptocurrencies with latest market data.
func (s *Client) GetLatestCryptocurrencyListing(options *CryptocurrencyListingsLatestOptions) ([]*api.PriceListing, error) {
	var params []string
	if options == nil {
		options = new(CryptocurrencyListingsLatestOptions)
	}
	if options.Limit != 0 {
		params = append(params, fmt.Sprintf("limit=%v", options.Limit))
	}
	url := fmt.Sprintf("%s/cryptocurrency/listings/latest?%s", s.baseURL, strings.Join(params, "&"))

	headers := make(map[string]interface{})
	headers[apiHeaderKeyName] = s.apiKey

	body, err := s.makeReq(url, headers)
	resp := new(api.Response)
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return nil, fmt.Errorf("JSON Error: [%s]. Response body: [%s]", err.Error(), string(body))
	}

	var listings []*api.PriceListing
	ifcs, ok := resp.Data.([]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for i := range ifcs {
		ifc := ifcs[i]
		listing := new(api.PriceListing)
		b, err := json.Marshal(ifc)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(b, listing)
		if err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}

	return listings, nil
}

// NewCoinMarketCapClient returns a new client instance for Market Cap Api.
func NewCoinMarketCapClient() (*Client, error) {
	return New(Config{}, BaseURL(coinMarketCapURL))
}
