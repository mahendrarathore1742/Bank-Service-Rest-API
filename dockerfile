FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /bank-service-rest-api

EXPOSE 3000

CMD [ "/bank-service-rest-api" ]