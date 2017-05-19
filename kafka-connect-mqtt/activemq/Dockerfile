FROM alpine
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

ARG ACTIVEMQ_VERSION="5.14.5"
RUN wget "http://apache.mirrors.ovh.net/ftp.apache.org/dist/activemq/${ACTIVEMQ_VERSION}/apache-activemq-${ACTIVEMQ_VERSION}-bin.tar.gz" \
            -O /activemq.tgz \
    && mkdir -p /opt/activemq \
    && tar -xzf /activemq.tgz --no-same-owner --strip-components=1 -C /opt/activemq \
    && rm -rf /activemq.tgz

ARG ACTIVEMQ_ADMIN="admin: admin, admin"
ARG ACTIVEMQ_USER="user: user, user"
RUN sed -e '/admin:.*/d' \
        -e '/user:.*/d' \
        -i /opt/activemq/conf/jetty-realm.properties \
    && echo "$ACTIVEMQ_ADMIN" >> /opt/activemq/conf/jetty-realm.properties \
    && echo "$ACTIVEMQ_USER" >> /opt/activemq/conf/jetty-realm.properties

#      tpp   ampq stomp mqtt webconsole
EXPOSE 61616 5672 61613 1883 8161
ENV PATH $PATH:/opt/activemq/bin

ENTRYPOINT [ "/opt/activemq/bin/activemq", "console" ]
