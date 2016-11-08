FROM landoop/fast-data-dev-ui:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

ENV GOPATH=/opt/go PATH=$PATH:/opt/go/bin
ADD main.go /opt/go/src/landoop/activemq-test/
RUN apk add --no-cache go git \
    && cd /opt/go/src/landoop/activemq-test/ \
    && go get -v \
    && go build -v \
    && apk del --no-cache git go pcre

ARG ACTIVEMQ_VERSION="5.14.1"
RUN wget "http://apache.mirrors.ovh.net/ftp.apache.org/dist/activemq/${ACTIVEMQ_VERSION}/apache-activemq-${ACTIVEMQ_VERSION}-bin.tar.gz" \
            -O /activemq.tgz \
    && mkdir -p /opt/activemq \
    && tar -xzf /activemq.tgz --no-same-owner --strip-components=1 -C /opt/activemq \
    && rm -rf /activemq.tgz

RUN rm -rf /extra-connect-jars/* \
    && ln -s /opt/activemq/activemq-all-${ACTIVEMQ_VERSION}.jar /extra-connect-jars/activemq-all-${ACTIVEMQ_VERSION}.jar
