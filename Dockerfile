FROM golang:1.23.3-alpine3.20

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

EXPOSE 9000
