FROM golang:1.22.0-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@v1.52.3

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./tmp/main ./cmd/main.go
 
EXPOSE 3000

CMD ["air"]