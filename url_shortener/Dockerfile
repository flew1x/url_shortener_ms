FROM golang:1.22.3-alpine

WORKDIR /app

COPY . /app

RUN go build -o url-shortener-ms /app/cmd/app

CMD ["./url-shortener-ms"]