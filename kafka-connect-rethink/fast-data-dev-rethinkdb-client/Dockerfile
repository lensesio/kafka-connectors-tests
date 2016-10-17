FROM landoop/fast-data-dev:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

ENV GOPATH=/opt/go PATH=$PATH:/opt/go/bin
ADD main.go /opt/go/src/landoop/rethinkdb-test/
RUN apk add --no-cache go git \
    && cd /opt/go/src/landoop/rethinkdb-test/ \
    && go get -v \
    && go build -v \
    && apk del --no-cache git go pcre
