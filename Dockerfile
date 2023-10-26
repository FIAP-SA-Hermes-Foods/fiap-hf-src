FROM golang:1.21

WORKDIR /hermes-foods/src/app

COPY . .

ENV GOPATH=/hermes-foods

RUN go build -ldflags "-w -s" -o bin/hermesfoods cmd/server/*.go

ENTRYPOINT ["./bin/hermesfoods"]

EXPOSE 8080
