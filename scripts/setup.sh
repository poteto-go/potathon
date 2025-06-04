#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

main () {
    helm upgrade --install ingress-nginx ingress-nginx \
        --repo https://kubernetes.github.io/ingress-nginx \
        --namespace ingress-nginx --create-namespace
}

main