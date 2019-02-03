#!/bin/bash

protoc -I ./api/price --go_out=plugins=grpc:./api/price ./api/price/*.proto
protoc -I ./api/ranking --go_out=plugins=grpc:./api/ranking ./api/ranking/*.proto

docker build -t top-coins/api -f build/api.Dockerfile .
docker build -t top-coins/price -f build/price.Dockerfile .
docker build -t top-coins/ranking -f build/ranking.Dockerfile .

kubectl apply -f deployments/api.yaml
kubectl apply -f deployments/price.yaml
kubectl apply -f deployments/ranking.yaml