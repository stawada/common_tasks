FROM node:18.17.1-alpine

RUN apk update && apk add \
    git

RUN mkdir front_app

WORKDIR /front_app

COPY . /front_app