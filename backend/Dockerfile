FROM golang:1.20.6-alpine

RUN apk update && apk add \
    git

RUN mkdir back_app

WORKDIR /back_app

COPY . /back_app

RUN apk --update add curl
