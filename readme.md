# kubernetes-diff-logger

This simple application is designed to watch Kubernetes objects and log diffs when they occur.  It is designed to log changes to Kubernetes objects in a clean way for storage and processing in [Loki](https://github.com/grafana/loki/).

## Example

```log
./kubernetes-diff-logger -namespace=default
{"timestamp":"2019-10-23T16:57:23Z","verb":"updated","type":"deployment","notes":"[Replicas: 1 != 2]", "name":"nginx"}}
{"timestamp":"2019-10-23T16:57:35Z","verb":"updated","type":"deployment","notes":"[Template.Spec.Containers.slice[0].Image: nginx != nginx:latest]", "name":"nginx"}}
```

See [Deployment](./deployment) for example yaml to deploy to Kubernetes. The example will monitor and log information about changes in all namespaces.

## Usage

```shell
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
- nameFilter: "*"
  type: "pod"
```
