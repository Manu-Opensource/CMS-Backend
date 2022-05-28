FROM golang:latest
WORKDIR /home/runner
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
EXPOSE 8080/tcp
COPY . .
CMD go run main.go
