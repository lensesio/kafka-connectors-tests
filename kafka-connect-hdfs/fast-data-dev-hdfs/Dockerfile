FROM landoop/fast-data-dev:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

ARG HADOOP_VERSION="2.7.3"
RUN wget https://www-us.apache.org/dist/hadoop/common/hadoop-${HADOOP_VERSION}/hadoop-${HADOOP_VERSION}.tar.gz \
        -O /hadoop.tgz \
    && mkdir /opt/hadoop \
    && tar xzf /hadoop.tgz --no-same-owner --strip-components=1 -C /opt/hadoop \
    && rm /hadoop.tgz \
    && ln -s /opt/hadoop /opt/hadoop-${HADOOP_VERSION}

ENV JAVA_HOME /usr/lib/jvm/java-1.8-openjdk/jre
ENV PATH ${PATH}:/opt/hadoop/bin:/opt/hadoop/sbin
ENV HADOOP_HOME /opt/hadoop

ADD hdfs.conf /etc/supervisord.d/
ADD core-site.xml hdfs-site.xml /opt/hadoop/etc/hadoop/

RUN hdfs namenode -format

#RUN rm -rf /extra-connect-jars/*
