FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD ["app", "./config.yaml"]

HEALTHCHECK --interval=10s --timeout=3s \
  CMD curl -f http://127.0.0.1/health || exit 1

EXPOSE 80
