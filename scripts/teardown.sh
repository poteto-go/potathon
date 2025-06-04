#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

teardown() {
    # delete ingress
    kubectl delete ingressclass nginx
    kubectl delete ing potathon-ingress

    # delete nodeport
    kubectl delete service potathon-nginx-np

    # delete loadbalancer
    kubectl delete service potathon-nginx-lb

    # delete nginx
    kubectl delete deployment potathon-nginx
    kubectl delete service potathon-nginx-service
    kubectl delete configmap nginx-conf

    # delete backend
    kubectl delete deployment potathon-backend
    kubectl delete service potathon-backend-service

    # delete redis
    kubectl delete deployment potathon-kvs
    kubectl delete service potathon-kvs-service
}

main() {
    teardown
}

main