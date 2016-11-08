# Connect UI Test

So, you wanna test Kafka Connect UI. Sure. Here is how:

Get into fast-data-dev-ui

    cd fast-data-dev-ui

Copy your Connect UI files:

    rm -rf kafka-connect-ui/*
    cp -r /path/to/my/connect/dev/* kafka-connect-ui/

Build fast-data-dev-ui:

    docker build -t landoop/fast-data-dev-ui .

Now go to a test. Elastic for example:

    cd ../kafka-connect-elastic

And run the test:

    coyote

Wait a bit for the containers to start and visit http://localost:3030/kafka-connect-ui

Once finished, remember to clean up your running containers:

    docker-compose down
