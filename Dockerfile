FROM golang:latest
WORKDIR /home/runner
EXPOSE 8080/tcp
COPY . .
CMD go run main.go
