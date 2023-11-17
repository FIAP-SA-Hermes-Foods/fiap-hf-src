FROM golang:1.21

WORKDIR /fiap-hf-src/src/app

COPY . .

ENV GOPATH=/fiap-hf-src

RUN go build -ldflags "-w -s" -o bin/hermesfoods cmd/server/*.go

ENTRYPOINT ["./bin/hermesfoods"]

EXPOSE 8080
