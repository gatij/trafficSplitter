FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o service-v1 .

EXPOSE 8081

CMD ["./service-v1"]
