FROM golang:1.14 AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o server .

EXPOSE 8877

CMD ["./server"]