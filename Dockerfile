FROM golang:1.20.5-bullseye as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod vendor

RUN go build -o /go/bin/app cmd/handler/main.go

FROM gcr.io/distroless/base-debian11
ENV TZ="Asia/Tokyo"
COPY --from=build /go/bin/app /
CMD ["/app"]