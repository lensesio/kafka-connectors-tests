# kafka-connect-elastic test #

This is the test for Landoop's Elastic Search connector (provided
by the Stream Reactor repo).

Most recent report: <http://coyote.landoop.com/connect/kafka-connect-elastic/>

---

## Usage

Just run coyote inside this directory.

    coyote

When finished, open `coyote.html`.

## Software

You need these programs to run the test:
- [Coyote](https://github.com/Landoop/coyote/releases)
- [Docker](https://docs.docker.com/engine/installation/)
- [Docker Compose](https://docs.docker.com/engine/installation/)

Everything else is set up automatically inside containers.

## Issues

ElasticSearch 5 needs system property `vm.max_map_count` at least at `262144`.
Some systems, like CentOS set it very low (`65530`). To fix it:

    sudo sysctl -w vm.max_map_count=262144
    echo -e "# for elasticsearch 5\nvm.max_map_count=262144" \
        | sudo tee /etc/sysctl.d/landoop-elasticsearch.conf
