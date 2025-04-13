FROM golang:1.23.2-alpine as builder

WORKDIR /app
COPY server.go .
RUN go build -o server server.go

FROM alpine:latest
COPY --from=builder /app/server /server
CMD ["/server"]
