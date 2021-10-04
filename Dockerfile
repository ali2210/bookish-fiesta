FROM golang:1.17.0-alpine3.14

ENV CGO_ENABLED=0

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o fiesta
 
CMD ["/app/fiesta"]