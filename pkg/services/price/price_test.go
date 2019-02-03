package main

import (
	"context"
	"log"
	"testing"

	pb "github.com/esanim/top-coins/api/price"
	"github.com/golang/protobuf/proto"
)

func TestGetPrices(t *testing.T) {
	err := setExternalClient()
	if err != nil {
		log.Fatalf("Failed to instantiate a crypto compare client: %v", err)
	}

	server := server{}
	req := &pb.PriceRequest{Limit: 1}
	listings, err := server.GetPrices(context.Background(), req)
	if err != nil {
		t.FailNow()
	}

	data, err := proto.Marshal(listings)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	if len(data) < 1 {
		t.FailNow()
	}
}
