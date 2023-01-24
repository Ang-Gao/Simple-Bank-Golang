#basic version
FROM golang:1.19.5-alpine3.16 AS builder
WORKDIR /app
COPY . ..
RUN go build -o main main.go

FROM alpine3.16
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]