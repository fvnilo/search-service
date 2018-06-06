FROM golang:1.10-alpine

RUN apk add --no-cache ca-certificates openssl git
RUN wget -O /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && \
  chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/nylo-andry/search-service
WORKDIR /go/src/github.com/nylo-andry/search-service

ADD . .

RUN dep ensure

RUN go build -o main .

CMD ["./main"]