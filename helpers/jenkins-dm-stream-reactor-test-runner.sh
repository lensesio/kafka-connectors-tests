#!/usr/bin/env bash

# This script can be used by jenkins to run the tests.
# We use it to avoid having to manage too many similar jenkins jobs.
# Instead of having to apply a small change to 20 jobs, we apply it here.
#
# The only argument it takes is the test directory name.

set -e

TEST_DIR="${TEST_DIR:-$1}"
TEST_VERSION="${TEST_VERSION:-$2}"
TEST_VERSION="${TEST_VERSION:latest}"
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
if [[ ! -d "$TEST_DIR" ]]; then
    echo "Test directory not found. The '\$1' argument is: $1."
    exit 255
fi
TEST_NAME="$(basename "$TEST_DIR")"

# Download latest Coyote
COYOTE=coyote-1.1_alpha_3-linux-amd64
wget -nc https://archive.landoop.com/tools/coyote/testing/$COYOTE
chmod +x $COYOTE

# Set path for jenkins
export PATH="$PATH:/usr/local/bin"
alias docker-compose="/usr/local/bin/docker-compose"

# cd into workdir, replace fast-data-dev version and run coyote
# then restore files via git
pushd "$TEST_DIR"
set +e
FILES_CHANGED="$(grep -rl 'landoop/fast-data-dev:latest' .)"
echo "$FILES_CHANGED" | xargs sed "s|landoop/fast-data-dev:latest|landoop/fast-data-dev:$TEST_VERSION|g" -i
sed -e "s/$TEST_NAME/$TEST_NAME $TEST_VERSION/" -i coyote.yml
"$WORKSPACE/$COYOTE"
EXITCODE="$?"
[[ ! -z "$FILES_CHANGED" ]] && echo "$FILES_CHANGED"  | xargs git checkout --
git checkout -- coyote.yml
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

# Create index page for test
rm -rf mainpage
mkdir -p mainpage
wget "https://raw.githubusercontent.com/Landoop/coyote-results-aggregator/master/index.html" -O mainpage/index.html
sed -e "s/TITLE/$TEST_DIR test/g" -i mainpage/index.html

# Disable error checking in case we run this locally.
set +e
cp latest.html "/usr/share/jenkins/public/stream-reactor/$TEST_VERSION/$TEST$(cat "$WORKSPACE"/status.txt).html"
set -e
