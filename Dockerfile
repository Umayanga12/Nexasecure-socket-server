# api1/Dockerfile (repeat for other APIs)
FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go build -o api1

EXPOSE 8080

CMD ["./api1"]
