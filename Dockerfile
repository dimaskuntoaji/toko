FROM golang:1.18

WORKDIR /app

COPY ./go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o toko1

EXPOSE 8080

CMD ["./toko1"]