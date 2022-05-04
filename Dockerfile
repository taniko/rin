FROM golang:1.18.1-bullseye as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod vendor

RUN go build -o /go/bin/app cmd/server/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /
CMD ["/app"]