package coins

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"

	api_coin "github.com/esanim/top-coins/api/coin"
	pb_price "github.com/esanim/top-coins/api/price"
	pb_ranking "github.com/esanim/top-coins/api/ranking"
	"github.com/esanim/top-coins/pkg/app"
	"google.golang.org/grpc"
)

// Handler process top coins endpoint
type Handler struct {
	app        *app.App
	rankingSvc *pb_ranking.RankingServiceClient
	pricingSvc *pb_price.PriceServiceClient
}

// NewCoinsHandler initialize VersionHandler object
func NewCoinsHandler(app *app.App) *Handler {

	// Connect to Rankings service
	connR, errR := grpc.Dial(":3001", grpc.WithInsecure())
	if errR != nil {
		log.Fatalf("Dial failed: %v", errR)
	}

	// Connect to Price service
	connP, errP := grpc.Dial(":3002", grpc.WithInsecure())
	if errP != nil {
		log.Fatalf("Dial failed: %v", errR)
	}

	rankingSvc := pb_ranking.NewRankingServiceClient(connR)
	pricingSvc := pb_price.NewPriceServiceClient(connP)

	return &Handler{app: app, pricingSvc: &pricingSvc, rankingSvc: &rankingSvc}
}

func (h *Handler) getCoinsFromRankingSvc(ch chan pb_ranking.RankingResponse, wg *sync.WaitGroup, limit int) {
	defer wg.Done()
	reqR := &pb_ranking.RankingRequest{Limit: uint64(limit)}
	if resR, errR := (*h.rankingSvc).GetRankings(context.Background(), reqR); errR == nil {
		ch <- *resR
	} else {
		log.Println(errR)
	}
}

func (h *Handler) getCoinsFromPriceSvc(ch chan pb_price.PriceResponse, wg *sync.WaitGroup, limit int) {
	defer wg.Done()
	reqP := &pb_price.PriceRequest{Limit: uint64(limit)}
	if resP, errP := (*h.pricingSvc).GetPrices(context.Background(), reqP); errP == nil {
		ch <- *resP
	} else {
		log.Println(errP)
	}
}

func (h *Handler) collectResult(ch1 chan pb_ranking.RankingResponse, ch2 chan pb_price.PriceResponse) (map[string]*api_coin.Coin, error) {
	coins := make(map[string]*api_coin.Coin, 10000)
	for i, s := range (<-ch1).Items {
		coins[s.Name] = &api_coin.Coin{Rank: int64(i), Symbol: s.Name}
	}
	for _, s := range (<-ch2).Items {
		if coins[s.Symbol] != nil {
			coins[s.Symbol].PriceUSD = s.PriceUSD
		}
	}
	return coins, nil
}

func (h *Handler) getCoins(limit int) []api_coin.Coin {
	priceChan := make(chan pb_price.PriceResponse, 1000)
	rankingChan := make(chan pb_ranking.RankingResponse, 1000)
	var wg sync.WaitGroup
	wg.Add(2)
	go h.getCoinsFromPriceSvc(priceChan, &wg, limit)
	go h.getCoinsFromRankingSvc(rankingChan, &wg, limit)
	wg.Wait()
	res, _ := h.collectResult(rankingChan, priceChan)

	resp := make([]api_coin.Coin, 0, len(res))
	for _, value := range res {
		resp = append(resp, *value)
	}
	return resp
}

// GetCoins retireves the coins data using merging 2 grpc calls.
func (h *Handler) GetCoins(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 200
	}
	if limit > 200 || limit < 1 {
		http.Error(w, "Limit should be greater than 0 and less than 200.", 500)
	}

	resp := h.getCoins(limit)

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Rank < resp[j].Rank
	})
	resp = resp[:limit]

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 500)
		h.app.Logger.Error().Err(err).Msg("Could not render response.")
	}
}
