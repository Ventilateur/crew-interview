FROM golang:1.18-alpine as builder

WORKDIR /crew

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o app


FROM alpine:3.16

COPY --from=builder /crew/app /crew/app
COPY --from=builder /crew/config/default.yaml /crew/config/default.yaml

ENTRYPOINT ["./crew/app"]
CMD ["server", "--config", "/crew/config/default.yaml"]
