package apiext

// CoinRanking is the listing structure
type CoinRanking struct {
	ID        string `json:"Id"`
	Name      string `json:"Name"`
	CoinName  string `json:"CoinName"`
	SortOrder string `json:"SortOrder"`
}

// CoinsRoot is the response structure
type CoinsRoot struct {
	Response string                 `json:"Response"`
	Data     map[string]CoinRanking `json:"Data"`
}
