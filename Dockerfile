FROM golang:1.14 as builder


WORKDIR /app
COPY . .

RUN go build -o main cmd/linkstore

FROM alpine:latest

COPY --from=builder /app/main main

CMD ["./main", "-code", "test"]