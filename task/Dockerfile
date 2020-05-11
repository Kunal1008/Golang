FROM golang:alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main server.go

EXPOSE 8080

CMD ["./main"]