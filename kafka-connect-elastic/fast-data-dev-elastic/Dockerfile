FROM landoop/fast-data-dev
MAINTAINER Marios Andreopoulos <marios@landoop.com>

WORKDIR /
RUN wget https://download.elastic.co/elasticsearch/release/org/elasticsearch/distribution/tar/elasticsearch/2.4.1/elasticsearch-2.4.1.tar.gz \
    && tar xf elasticsearch-2.4.1.tar.gz \
    && mv /elasticsearch-2.4.1/lib/*.jar /connectors/ \
    && rm -rf /elasticsearch-2.4.1*
