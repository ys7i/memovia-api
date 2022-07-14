FROM golang:1.18-alpine

WORKDIR /go

COPY ./src /go/
COPY ./go.mod /go/
COPY ./go.sum /go/

ENV GOPATH=

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go get && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
