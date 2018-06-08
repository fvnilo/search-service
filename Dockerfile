FROM golang:1.10-alpine

RUN mkdir -p /go/src/github.com/nylo-andry/search-service
WORKDIR /go/src/github.com/nylo-andry/search-service

ADD . .

RUN go build -o main .

CMD ["sh", "start.sh"]