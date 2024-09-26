# Dockerfile
FROM golang:1.23

RUN go install github.com/air-verse/air@latest

WORKDIR /app

EXPOSE 5001

CMD ["./air_dev.sh"]