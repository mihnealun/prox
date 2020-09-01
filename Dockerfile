FROM golang:1.13-alpine

RUN apk add --purge --no-cache --update inotify-tools git wv libc-dev gcc

RUN mkdir -p /go/src/github.com/mihnealun/prox] && go get -u github.com/go-delve/delve/cmd/dlv

WORKDIR /go/src/github.com/mihnealun/prox

CMD ./docker/autobuild.sh;