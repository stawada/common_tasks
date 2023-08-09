FROM golang:1.20.6-alpine

RUN apk update && apk add \
    git

WORKDIR /app

COPY ./backend /app


