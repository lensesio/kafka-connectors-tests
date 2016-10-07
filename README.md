# Connector Tests #

Basic testing of various connectors using docker and coyote.



## Prepare Test Server ##

For now we use Cloudera03 server. This will hopefully not be a permanent setup
because it gives too much rights to jenkins.
Thus instead of automating the needed steps via ansible, we will document them
here.

Permit Jenkins user to run docker and generally have some root rights by addding
it to root group:

    sudo usermod -aG docker jenkins

Install docker compose via an official release since it isn't yet available in
centos.
Visit [compose github release page](https://github.com/docker/compose/releases)
for the latest release. In general, you will run something like this:

    sudo su
    curl -L https://github.com/docker/compose/releases/download/1.8.1/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose

## Running in Jenkins

We have many tests which we want to run separately and in parallel. This leads
to one jenkins job per test design, which unfortunately is time-consuming to
maintain, as for each little change, we would have to update tens of jenkins'
jobs.

For this we created the `helpers` directory and there we store scripts we use
to run the tests from jenkins, thus moving the better part of the run logic to
this repo.

To run a test from Jenkins, you would run something like:

    helpers/jenkins-test-runner.sh kafka-connect-redis

This runs the tests and creates such files (as status.txt and exitcode) that
the jenkins job can use to set the build name, exit status, etc.
