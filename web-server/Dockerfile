FROM golang:alpine

COPY . /app

WORKDIR /app/cmd/

RUN go build -v

EXPOSE 8080

ENTRYPOINT ["/app/cmd/cmd"]
