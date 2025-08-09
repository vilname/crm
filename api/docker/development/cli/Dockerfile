# Стадия сборки
FROM golang:1.24.6-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
