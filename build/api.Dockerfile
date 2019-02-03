FROM grpc/go:latest

ENV GO111MODULE=on

WORKDIR /go/src/github.com/esanim/top-coins
COPY . .

RUN go get -v ./...
RUN go install -v ./...

ENV PORT 3000
EXPOSE ${PORT}

ENTRYPOINT ["cmd/main"]