package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	pb "github.com/esanim/top-coins/api/ranking"
	ext "github.com/esanim/top-coins/pkg/services/external"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var client *ext.Client

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	err = setExternalClient()
	if err != nil {
		log.Fatalf("Failed to instantiate a crypto compare client: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRankingServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func setExternalClient() error {
	c, err := ext.NewCryptoCompareClient()
	if err != nil {
		return err
	}
	client = c
	return nil
}

// GetRankings retrieves a list of coins data with the sort order info.
func (s *server) GetRankings(ctx context.Context, r *pb.RankingRequest) (*pb.RankingResponse, error) {
	coinList, err := client.ListCoins()
	if err != nil {
		return nil, err
	}

	var coins []*pb.Ranking
	for _, val := range coinList {
		sortOrder, _ := strconv.Atoi(val.SortOrder)
		res := &pb.Ranking{Id: val.ID, Name: val.Name, SortOrder: uint64(sortOrder)}
		coins = append(coins, res)
	}

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
		return nil, err
	}
	return &pb.RankingResponse{Items: coins}, nil
}
