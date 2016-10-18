Here we have two java client app to help us test the connector.

The `create-table` app create the `connect_test` table into Kudu. The
`read-table` app reads all entries from the `connect_test` table in Kudu.

To build and run, do the following in each application's directory:

    $ mvn package
    $ java -jar target/kudu-java-sample-1.0-SNAPSHOT.jar

The default host is localhost. To run with a different host:

    $ java -DkuduMaster=<HOST> -jar target/kudu-java-sample-1.0-SNAPSHOT.jar

---

The tests are copied from Cloudera's Kudu Examples: https://github.com/cloudera/kudu-examples/tree/master/java/java-sample
