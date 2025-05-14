#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

log() {
  echo "$@" > /dev/tty
}

die() {
  log "$@"
  exit 1
}
