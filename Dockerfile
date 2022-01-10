# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o /covid-19-summary

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /covid-19-summary ./
EXPOSE 8080
CMD ["./covid-19-summary"]