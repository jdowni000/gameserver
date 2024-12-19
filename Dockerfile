FROM golang:alpine

COPY . /app

WORKDIR /app/cmd/web-server

RUN go build -v

EXPOSE 8080

ENTRYPOINT ["/app/cmd/web-server/web-server"]
