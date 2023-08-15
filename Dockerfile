FROM golang:1.20.6-alpine

RUN apk update && apk add \
    git

RUN mkdir back_app

RUN mkdir front_app

WORKDIR /back_app

COPY ./backend /back_app

COPY ./vue-test-project /front_app