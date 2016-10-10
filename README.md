# Kafka Connectors Tests

An independent set of tests for various Kafka Connect connectors.

---

## Introduction

We setup and test various connectors in a pragmatic environment. That is we
spawn at least a broker, a zookeeper instance, a schema registry, a connect
distributed instance and any other software needed (e.g elasticsearch, redis,
cassandra) in docker containers and then perform tests using standard tools.
This practice permits us to verify that a connector does work, as well as
provide a basic example of how to setup and test it. Advanced tests verify how
the connector performs in special cases.

To achieve this we use our in-house developed —open source— tools coupled with
docker-compose. The main testing tool is
[Coyote](https://github.com/Landoop/coyote), which takes yml files describing
the test process and performs each step logging output, errors and other
information. Our
[fast-data-dev](https://hub.docker.com/r/landoop/fast-data-dev/) docker image
is used as a reference Confluent Platform installation.

## Run the tests

To run the tests on your computer you need coyote, docker and docker-compose.
You can grab coyote from our
[release page](https://github.com/Landoop/coyote/releases) or built it yourself
via `go get github.com/landoop/coyote`. For the installation of docker and
docker-compose we will have to refer you to
[docker's](https://docs.docker.com/engine/installation/) and
[docker-compose's](https://docs.docker.com/engine/installation/) documentation.

Once you install all the tools, just enter into a test directory and run:

    coyote

Wait a few minutes for coyote to finish and it will produce a `coyote.html` file
with the test's report.
Coyote uses its exit code to indicate the numbers of tests that failed, thus an
error code from coyote doesn't usually show a problem into coyote itself.
