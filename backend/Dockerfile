FROM golang:1.20.6-alpine

RUN apk update && apk add \
    git

RUN mkdir back_app

COPY . /back_app

WORKDIR /back_app/api_server

RUN apk --update add curl

CMD ["go", "run", "api.go", "main.go"]
