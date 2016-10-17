This docker image servers two purposes.

1.
We add the activemq jar into connect's classpath, so the JMS connector can work
with ActiveMQ.

2.
We add and compile a small golang program that connects to the testing
ActiveMQ's topic and reads the messages we send, in order for us to verify that
the connector indeed works (and not only claims that it works).
