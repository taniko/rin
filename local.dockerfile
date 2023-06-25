FROM golang:1.20.4-bullseye
ENV TZ="Asia/Tokyo"
WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

ADD . /go/src/app
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod vendor
CMD ["air"]
