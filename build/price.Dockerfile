FROM grpc/go:latest

WORKDIR /go/src/github.com/esanim/top-coins/pkg/services/price
COPY api/price ../../../api/price
COPY pkg/services/external ../external
COPY pkg/services/price .

RUN go get -v ./...
RUN go build -v ./...

ENV PORT 3002
EXPOSE ${PORT}

ENTRYPOINT [ "price" ]