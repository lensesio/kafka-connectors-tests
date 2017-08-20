FROM landoop/fast-data-dev:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

RUN apk add --no-cache \
        bash \
        wget \
        tar \
        gzip \
        ca-certificates \
        openjdk8-jre-base \
        supervisor \
    && echo "progress = dot:giga" | tee /etc/wgetrc \
    && mkdir -p /opt

ARG HBASE_VERSION="1.3.1"
RUN wget "https://www-eu.apache.org/dist/hbase/1.3.1/hbase-1.3.1-bin.tar.gz" \
    && tar xzf "/hbase-${HBASE_VERSION}-bin.tar.gz" --no-same-owner -C /opt/

RUN rm -rf /extra-connect-jars/*

ADD hbase-site.xml /opt/hbase-${HBASE_VERSION}/conf/

ENV JAVA_HOME=/usr/lib/jvm/default-jvm
ENV PATH="$PATH:/opt/hbase-${HBASE_VERSION}/bin"
