FROM cassandra:3.0
MAINTAINER Marios Andreopoulos <marios@landoop.com>

RUN sed -e 's/authenticator: AllowAllAuthenticator/authenticator: PasswordAuthenticator/' \
        -i /etc/cassandra/cassandra.yaml
