# metrics-collector

MetricsCollector is responsible of collecting metrics and informations about resources in a Kubernetes cluster.


## Developping

### Environment

We use [microk8s](https://microk8s.io/) to run a kubernetes cluster for our dev environment.

Once installed you need to install some addons:
```sh
microk8s.enable metrics-server storage dns dashboard ingress helm rbac prometheus
```