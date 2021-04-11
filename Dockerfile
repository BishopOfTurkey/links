FROM golang:1.14 as builder


WORKDIR /app
COPY . .

RUN go build -o main cmd/linkstore.go

FROM alpine:latest

COPY --from=builder /app/main main
COPY static static

CMD ["./main", "-code", "test"]