FROM landoop/fast-data-dev
MAINTAINER Marios Andreopoulos <marios@landoop.com>

ADD kafka-connect-ui/ /var/www/kafka-connect-ui

RUN sed -r -e 's|https://schema-registry.demo.landoop.com|/api/schema-registry|' \
           -e 's|http://localhost:8081|/api/schema-registry|' \
           -e 's|https://kafka-rest-proxy.demo.landoop.com|/api/kafka-rest-proxy|' \
           -e 's|http://localhost:8082|/api/kafka-rest-proxy|' \
           -e 's|https://kafka-connect.demo.landoop.com|/api/kafka-connect|' \
           -e 's|http://localhost:8083|/api/kafka-connect|' \
           -i /var/www/kafka-connect-ui/combined.js
