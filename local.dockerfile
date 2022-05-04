FROM golang:1.18.1-bullseye

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@v1.29.0


ADD . /go/src/app
RUN go mod vendor
CMD ["air"]
