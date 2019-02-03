FROM grpc/go:latest

WORKDIR /go/src/github.com/esanim/top-coins/pkg/services/ranking
COPY api/ranking ../../../api/ranking
COPY pkg/services/external ../external
COPY pkg/services/ranking .

RUN go get -v ./...
RUN go build -v ./...

ENV PORT 3001
EXPOSE ${PORT}

ENTRYPOINT [ "ranking" ]