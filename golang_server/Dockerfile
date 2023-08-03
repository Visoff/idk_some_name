FROM golang:1.19

WORKDIR /var/golang/server

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main

EXPOSE 8080

CMD [ "./main" ]