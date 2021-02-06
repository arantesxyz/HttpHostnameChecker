FROM golang:1.15.8-alpine3.13

WORKDIR /usr/src/app

ENV PORT 8080
ENV HOST 0.0.0.0

COPY . .

RUN go build -o app .

CMD ./app