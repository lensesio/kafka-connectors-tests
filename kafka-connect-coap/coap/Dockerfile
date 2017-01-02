FROM landoop/fast-data-dev:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

ENV GOPATH=/opt/go PATH=$PATH:/opt/go/bin
RUN apk add --no-cache go git \
    && go get -v github.com/dustin/go-coap \
    && go get -v github.com/dustin/go-coap/example/server \
    && sed -e 's|some/path|path|g' -i /opt/go/src/github.com/dustin/go-coap/example/client/goap_client.go \
    && go get -v github.com/dustin/go-coap/example/client \
    && ls /opt/go/bin \
    && apk del --no-cache git go pcre

EXPOSE 5683
ENTRYPOINT [ "/opt/go/bin/server" ]
