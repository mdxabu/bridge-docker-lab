FROM golang:1.21-alpine as builder

WORKDIR /app
COPY server.go .

RUN go build -o server server.go

FROM alpine:latest
COPY --from=builder /app/server /server

ENTRYPOINT ["/server"]
