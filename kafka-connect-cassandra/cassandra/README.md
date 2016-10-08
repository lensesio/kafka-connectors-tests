This is a simple docker image which takes the official cassandra
image and enables basic authentication via a sed in the `cassandra.yaml`
configuration file.

It doesn't set anything else, thus the default credentials are kept:
`cassandra`/`cassandra`.
