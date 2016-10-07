#!/usr/bin/env bash

# This script can be used by jenkins to run the tests.
# We use it to avoid having to manage too many similar jenkins jobs.
# Instead of having to apply a small change to 20 jobs, we apply it here.
#
# The only argument it takes is the test directory name.

set -e

TEST_DIR="$1"
RESULTS_DIR="results"

# Set current dir (which should be repo's top dir) to workspace so we can test
# outside jenkins
WORKSPACE="${WORKSPACE:-$(pwd)}"

# Remove old files it they exist:
rm -f "$WORKSPACE"/status.txt
rm -f "$WORKSPACE"/exitcode
rm -f "$WORKSPACE"/latest.html
rm -f "$WORKSPACE"/index.html

# Check if test directory exists
if [[ ! -d "$1" ]]; then
    echo "Test directory not found. The '\$1' argument is: $1."
    exit 255
fi

# Download latest Coyote
COYOTE=coyote-1.0-amd64
wget -nc https://github.com/Landoop/coyote/releases/download/v1.0/$COYOTE
chmod +x $COYOTE

# Set path for jenkins
export PATH="$PATH:/usr/local/bin"
alias docker-compose="/usr/local/bin/docker-compose"

# cd into workdir and run coyote
pushd "$TEST_DIR"
set +e
"$WORKSPACE/$COYOTE"
EXITCODE="$?"
set -e
popd

# Store error number for jenkins build name
if [[ $EXITCODE -eq 0 ]]; then
    echo "" > "$WORKSPACE"/status.txt
elif [[ $EXITCODE -eq 1 ]]; then
    echo "_${EXITCODE}_err" > "$WORKSPACE"/status.txt
else
    echo "_${EXITCODE}_errs" > "$WORKSPACE"/status.txt
fi

# Store exitcode to use for output
echo "$EXITCODE" > "$WORKSPACE"/exitcode


# Copy test results to results directory
mkdir -p "$RESULTS_DIR"
DATE="$(date '+%Y%m%d-%H%M')"
cp "$TEST_DIR"/coyote.html "$RESULTS_DIR"/"$(basename ${TEST_DIR})-${DATE}$(cat status.txt).html"
rm -f latest.html
mv "$TEST_DIR"/coyote.html latest.html
ln -s latest.html index.html
