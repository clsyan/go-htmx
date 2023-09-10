FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash

RUN apt-get update

RUN apt-get install -y migrate

RUN go build -o main

EXPOSE 8080

CMD ["air"]