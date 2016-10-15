# Considerations #

1.
For now we use the `default` retention policy name because stream-reactor's
stable version has it hardcoded. In the future, the test should be adjusted with
a custom retention policy name.

2.
In order to make the connector work, we have to remove the rest of the
connectors from the classpath. It is up to us to decide whether this is a bug or
acceptable practice.
