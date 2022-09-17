FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN cd /app/demo/kafka/producer && go build -o /producer
RUN cd /app/demo/kafka/consumer && go build -o /consumer
RUN apk add vim curl bash