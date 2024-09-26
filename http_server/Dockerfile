# Dockerfile
FROM golang:1.23

RUN go install github.com/air-verse/air@latest

WORKDIR /app

EXPOSE 8080

CMD ["./air_dev.sh"]