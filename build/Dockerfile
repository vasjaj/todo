FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD ["app", "./config.yaml"]

EXPOSE 80
