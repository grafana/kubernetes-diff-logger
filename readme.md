# kubernetes-diff-logger

This simple application is designed to watch Kubernetes objects and log diffs when they occur. It is designed to log changes to Kubernetes objects in a clean way for storage and processing in [Loki](https://github.com/grafana/loki/).

## Example

```shell
./kubernetes-diff-logger -namespace=default
{"timestamp":"2019-10-23T16:57:23Z","verb":"updated","type":"deployment","notes":"[Replicas: 1 != 2]", "name":"nginx"}}
{"timestamp":"2019-10-23T16:57:35Z","verb":"updated","type":"deployment","notes":"[Template.Spec.Containers.slice[0].Image: nginx != nginx:latest]", "name":"nginx"}}
```

See [Deployment](./deployment) for example yaml to deploy to Kubernetes. The example will monitor and log information about changes in all namespaces.

## Usage

```text
Usage of ./kubernetes-diff-logger:
  -config string
     Path to config file. Required.
  -kubeconfig string
     Path to a kubeconfig. Only required if out-of-cluster.
  -log-added
     Log when deployments are added.
  -log-deleted
     Log when deployments are deleted.
  -master string
     The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.
  -namespace string
     Filter updates by namespace. Leave empty to watch all.
  -resync duration
     Periodic interval in which to force resync objects. (default 30s)
```

## Config File

```yaml
differs:
- nameFilter: "*"
  type: "deployment"
- nameFilter: "*"
  type: "statefulset"
- nameFilter: "*"
  type: "daemonset"
```

## Contibuting

`kubernetes-diff-logger` is currently built using Go v1.20.

If you install Go via [`asdf`](https://asdf-vm.com), you should set the `ASDF_GOLANG_MOD_VERSION_ENABLED` environment variable to `true` for forward compatibility.

### Building

```shell
go mod download
CGO_ENABLED=0 GOOS=linux go build -a -o app .
docker build .
```

### Testing

Testing is currently limited to manual interaction with the cluster.

In one shell:

```shell
cd test
./test.sh
```

In another shell:

```shell
cd test
kubectl create -f daemonset.yaml
kubectl create -f deployment.yaml
kubectl create -f statefulset.yaml
kubectl delete -f statefulset.yaml
kubectl delete -f deployment.yaml
kubectl delete -f daemonset.yaml
```

When you're done, press `Ctrl-C` in the first shell.
