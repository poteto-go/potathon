# potathon

## K8s

### Setup

- install kubectl

  https://kubernetes.io/ja/docs/tasks/tools/install-kubectl-linux/

- image build

```bash
docker build --pull --rm -t potathon-calculator:latest ./calculator
docker tag potathon-calculator <user>s/potathon-calculator:latest
docker push <user>/potathon-calculator:latest
```

### Calculator

Apply 10 replicaSet

```bash
kubectl apply -f ./k8s/calculator.yaml
```

### Teardown

```bash
go run main.go kube teardown
```

### How to

```bash
docker compose up
```
