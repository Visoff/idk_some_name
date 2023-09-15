FROM golang:1.20

WORKDIR /app/

RUN apt-get update && apt-get install -y git

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/bin/git_server

EXPOSE 8080
CMD ["git_server"]