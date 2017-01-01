FROM landoop/fast-data-dev:latest
MAINTAINER Marios Andreopoulos <marios@landoop.com>

WORKDIR /

# # This would build the client inside the container but it would be slow to do frequently.
# ENV SBT_OPTS="-Xms512M -Xmx1536M -Xss1M -XX:+CMSClassUnloadingEnabled"
# RUN apk --no-cache add git \
#     && git clone https://github.com/Landoop/hazelcast-client.git /hazelcast-client \
#     && wget https://repo.typesafe.com/typesafe/ivy-releases/org.scala-sbt/sbt-launch/0.13.13/sbt-launch.jar \
#     && cd /hazelcast-client \
#     && java $SBT_OPTS -jar /sbt-launch.jar assembly \
#     && rm -rf /sbt-launch.jar /root/.sbt /root/.ivy2

# ENTRYPOINT ["/usr/bin/java", "-jar", "/hazelcast-client/target/scala-2.11/hazelcast-client.jar"]

RUN wget https://github.com/Landoop/hazelcast-client/releases/download/0.3/hazelcast-client-0.3.jar \
         -O /hazelcast-client.jar

ENTRYPOINT ["/usr/bin/java", "-jar", "/hazelcast-client.jar" ]
