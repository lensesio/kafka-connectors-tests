FROM centos:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

WORKDIR /

RUN yum install -y epel-release wget \
    && wget https://archive.cloudera.com/beta/kudu/redhat/7/x86_64/kudu/cloudera-kudu.repo \
         -O /etc/yum.repos.d/cloudera-kudu.repo \
    && rpm --import \
           https://archive.cloudera.com/beta/kudu/redhat/7/x86_64/kudu/RPM-GPG-KEY-cloudera \
    && yum install -y \
                   supervisor \
                   kudu \
                   kudu-master \
                   kudu-tserver \
                   kudu-client0 \
                   maven \
    && yum clean all -y

COPY test-apps /test-apps
RUN cd /test-apps/create-table \
    && mvn package \
    && cd /test-apps/read-table \
    && mvn package

# RUN wget https://archive.cloudera.com/cdh5/redhat/7/x86_64/cdh/cloudera-cdh5.repo \
#          -O /etc/yum.repos.d/cloudera-cdh5.repo \
#     && rpm --import \
#            https://archive.cloudera.com/cdh5/redhat/7/x86_64/cdh/RPM-GPG-KEY-cloudera \
#     && wget https://archive.cloudera.com/beta/impala-kudu/redhat/7/x86_64/impala-kudu/cloudera-impala-kudu.repo \
#          -O /etc/yum.repos.d/cloudera-impala-kudu.repo \
#     && rpm --import \
#            https://archive.cloudera.com/beta/impala-kudu/redhat/7/x86_64/impala-kudu/RPM-GPG-KEY-cloudera \
#     && yum install -y \
#                    java-1.8.0-openjdk-headless \
#                    impala-kudu \
#                    impala-kudu-shell

EXPOSE 7050 8050 7051 8051

ADD supervisord.conf /etc/supervisord.conf

ENTRYPOINT ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
