# kubernetes-diff-logger

This simple application is designed to watch Kubernetes objects and log diffs when they occur.  

## example

```
./kubernetes-diff-logger -namespace=default
updated : nginx [Replicas: 1 != 2]
updated : nginx [Template.Spec.Containers.slice[0].Image: nginx != nginx:latest]
```

## usage

```
Usage of ./kubernetes-diff-logger:
  -kubeconfig string
    	Path to a kubeconfig. Only required if out-of-cluster.
  -log-added
    	Log when deployments are added.
  -log-deleted
    	Log when deployments are deleted.
  -master string
    	The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.
  -name-filter string
    	Glob based filter.  Only deployments matching will be processed. (default "*")
  -namespace string
    	Filter updates by namespace.  Leave empty to watch all.
  -resync duration
    	Periodic interval in which to force resync objects. (default 30s)
```