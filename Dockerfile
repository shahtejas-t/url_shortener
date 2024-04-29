FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o url_shortener ./cmd/url_shortener

CMD ["./url_shortener"]
