FROM landoop/fast-data-dev:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

RUN apk add --no-cache git openjdk8 maven \
    && git clone https://github.com/couchbase/kafka-connect-couchbase.git \
    && cd kafka-connect-couchbase \
    && mvn package \
    && cp target/kafka-connect-couchbase-*-package/share/java/kafka-connect-couchbase/* /connectors \
    && cd / && rm -rf kafka-connect-couchbase \
    && apk del --no-cache git maven

