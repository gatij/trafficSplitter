FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o service-v2 .

EXPOSE 8082

CMD ["./service-v2"]
