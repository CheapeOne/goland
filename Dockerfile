FROM golang:latest

WORKDIR /goland

COPY go.mod go.sum ./

RUN go mod download

COPY . .

