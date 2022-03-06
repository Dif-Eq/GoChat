# syntax=docker/dockerfile:1
FROM golang:1.17-alpine AS dependencies
WORKDIR /usr/local/src
COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM dependencies AS build
WORKDIR /usr/local/src
COPY . .
RUN go mod tidy
# RUN printenv
RUN go build -o /app ./cmd/main.go

FROM alpine:latest
COPY --from=build /app /app
EXPOSE 8080
CMD [ "/app"]
