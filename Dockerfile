# syntax=docker/dockerfile:1

# Билд контейнер для сборки Go приложения
FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main_car_service ./cmd/main.go

CMD ["ls"]

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main_car_service /root/main_car_service
COPY --from=builder /app/config/ /root/config/

EXPOSE 8080

CMD ["/root/main_car_service"]