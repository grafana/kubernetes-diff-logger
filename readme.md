# kubernetes-diff-logger

This simple application is designed to watch Kubernetes objects and log diffs when they occur.  

## example

```
./kubernetes-diff-logger -namespace=default
2019-10-23T14:40:06Z updated : nginx (deployment) [Replicas: 1 != 2]
2019-10-23T14:40:06Z updated : nginx (deployment) [Template.Spec.Containers.slice[0].Image: nginx != nginx:latest]
```

## usage

```
Usage of ./kubernetes-diff-logger:
  -config string
    	Path to config file.  Required.
  -kubeconfig string
    	Path to a kubeconfig. Only required if out-of-cluster.
  -log-added
    	Log when deployments are added.
  -log-deleted
    	Log when deployments are deleted.
  -master string
    	The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.
  -namespace string
    	Filter updates by namespace.  Leave empty to watch all.
  -resync duration
    	Periodic interval in which to force resync objects. (default 30s)
```

## config file

```
differs:
- nameFilter: "*"
  type: "deployment"
- nameFilter: "*"
  type: "statefulset"
- nameFilter: "*"
  type: "daemonset"
```