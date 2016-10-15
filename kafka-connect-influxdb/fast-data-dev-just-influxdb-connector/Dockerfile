FROM landoop/fast-data-dev:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

RUN rm -rf /extra-connect-jars/* \
    && mv /opt/confluent-3.0.1/share/java/kafka-connect-influxdb / \
    && rm -rf /opt/confluent-3.0.1/share/java/kafka-connect-* \
    && mv /kafka-connect-influxdb /opt/confluent-3.0.1/share/java/
