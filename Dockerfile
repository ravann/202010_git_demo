FROM golang:1.15.3-buster

LABEL version=1
LABEL author=ravan.nannapaneni@gmail.com

COPY ./src src/
COPY entrypoint.sh /

RUN chmod 755 /entrypoint.sh

RUN go get -u github.com/gorilla/mux

ENTRYPOINT /entrypoint.sh
