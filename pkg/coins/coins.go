package coins

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/esanim/top-coins/pkg/app"
)

// CoinsHandler process top coins endpoint
type CoinsHandler struct {
	app *app.App
}

// NewCoinsHandler initialize VersionHandler object
func NewCoinsHandler(app *app.App) *CoinsHandler {
	return &CoinsHandler{app: app}
}

// Rank,	Symbol,	Price USD
type coinResp struct {
	Rank     int64   `json:"rank"`
	Symbol   string  `json:"symbol"`
	PriceUSD float64 `json:"priceusd"`
}

type coinsResp struct {
	Data []coinResp `json:"data"`
}

func (h *CoinsHandler) GetCoins(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 200
	}
	if limit > 200 || limit < 1 {
		http.Error(w, "Limit should be greater than 0 and less than 200.", 500)
	}

	resp := coinsResp{
		Data: []coinResp{
			{1, "BTC", 6634.41},
			{2, "ETH", 370.237},
			{3, "XRP", 0.471636},
		},
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 500)
		h.app.Logger.Error().Err(err).Msg("Could not render version")
	}
}
