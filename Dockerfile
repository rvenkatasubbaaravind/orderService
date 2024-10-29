# server/Dockerfile
FROM golang:1.23

WORKDIR /app

COPY . .

RUN GOARCH=arm64 go build -o main .
EXPOSE 8080
EXPOSE 5000

CMD ["/bin/sh", "-c", "./main"]