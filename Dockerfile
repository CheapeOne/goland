FROM golang:latest

WORKDIR /goland

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# hot reloading
RUN go get -u github.com/githubnemo/CompileDaemon
CMD CompileDaemon -command="./goland"
