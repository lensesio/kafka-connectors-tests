# kafka-connect-couchbase test #

Most recent report: <http://coyote.landoop.com/connect/kafka-connect-couchbase/>

This test builds kafka-connect-couchbase (source and sink) from master and tests
it. Apart from basic testing we do a bit of CDC testing.

For the sink we update entries in the sink topic and expect to see updated
documents in the bucket. We can't test delete yet because we can't sent messages
with null value and avro key to a topicet through the command line keys.

For the source, we update a document in couchbase and delete another, then read
the source topic to detect these events.

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
