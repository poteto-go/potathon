#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

deploy() {
    # deploy redis
    kubectl apply -f k8s/redis/deploy.yaml
    kubectl apply -f k8s/redis/service.yaml

    # deploy backend
    kubectl apply -f k8s/backend/deploy.yaml
    kubectl apply -f k8s/backend/service.yaml

    # deploy nginx
    kubectl create configmap nginx-conf --from-file=k8s/nginx/nginx.conf
    kubectl apply -f k8s/nginx/deploy.yaml
    kubectl apply -f k8s/nginx/service.yaml

    # deploy loadbalancer
    kubectl apply -f k8s/loadbalancer/service.yaml

    local url=$(minikube service potathon-nginx-lb --url)
    echo service on "$url"
}

main() {
    deploy
}

main
