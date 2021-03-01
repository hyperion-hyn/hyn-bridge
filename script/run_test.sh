#!/bin/bash

SCRIPT_PARENT_PATH=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
source "$SCRIPT_PARENT_PATH/ganache.sh"
PROJECT_PATH="$(dirname "$SCRIPT_PARENT_PATH")"

bootstrap_ganache $PROJECT_PATH

# Run test
truffle compile --all
truffle test --network development

shut_down_ganache $PROJECT_PATH
