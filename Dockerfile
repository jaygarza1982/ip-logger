FROM golang:1.22-alpine

WORKDIR /app

# Dependencies
COPY ./go.mod .

COPY . .

RUN go build -o /usr/local/bin/ipv4logger .

CMD [ "ipv4logger" ]