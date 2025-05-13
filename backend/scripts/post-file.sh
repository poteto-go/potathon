#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

source scripts/common.sh

request () {
  local fileName="$1"

  log $fileName

  curl -X POST -F file=@${fileName} -F file-name=${fileName} http://localhost:3000/upload/input
}

main() {
  if [ $# -ne 1 ]; then
    die "arg size one, give me the filename"
  fi

  request "$1"
}

main "$@"