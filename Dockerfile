FROM golang:latest

WORKDIR /goland

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# hot reloading
RUN ["go", "get", "github.com/pilu/fresh"]
CMD [ "fresh" ]
