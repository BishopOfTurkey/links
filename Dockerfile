FROM golang:1.14 as builder


WORKDIR /app
COPY . .

RUN go build cmd/linkstore -o main

FROM alpine:latest

COPY --from builder /app/main main
