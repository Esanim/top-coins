package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/esanim/top-coins/api/price"
	ext "github.com/esanim/top-coins/pkg/services/external"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var client *ext.Client

func main() {
	lis, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	err = setExternalClient()
	if err != nil {
		log.Fatalf("Failed to instantiate a crypto compare client: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPriceServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func setExternalClient() error {
	c, err := ext.NewCoinMarketCapClient()
	if err != nil {
		return err
	}
	client = c
	return nil
}

// GetPrices retrieves a list of coins data with the usd price info.
func (s *server) GetPrices(ctx context.Context, r *pb.PriceRequest) (*pb.PriceResponse, error) {
	coinList, err := client.GetLatestCryptocurrencyListing(&ext.CryptocurrencyListingsLatestOptions{})

	var coins []*pb.Price
	for _, val := range coinList {
		q := val.Quote["USD"]
		res := &pb.Price{Id: val.ID, Name: val.Name, Symbol: val.Symbol, PriceUSD: float64(q.Price)}
		coins = append(coins, res)
	}

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
		return nil, err
	}
	return &pb.PriceResponse{Items: coins}, nil
}
