# Stage 1: Build the Go app
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Stage 2: Run the Go app
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
