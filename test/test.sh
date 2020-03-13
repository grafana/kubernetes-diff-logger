#!/usr/bin/env bash
K3D_NAME=k8s-diff-logger
CONFIG_PATH="k3d get-kubeconfig --name=$K3D_NAME"

# requires k3d
k3d create --name $K3D_NAME --api-port=6444

echo Building executable... 
go build ../

echo -n "Ensuring k3d is running..."
while true; do
  k3d list 2>&1 | grep ".*$K3D_NAME.*running" >/dev/null && echo "done" && break \
    || (echo -n . && sleep 1)
done

echo -n "Getting kubeconfig..."
while true; do
  eval $CONFIG_PATH 2>&1 | grep "$K3D_NAME/kubeconfig.yaml" >/dev/null && echo done && break \
    || (echo -n . && sleep 1)
done
echo Config is available at $(eval $CONFIG_PATH)

echo Running kubernetes-diff-logger...
./kubernetes-diff-logger -kubeconfig=$(eval $CONFIG_PATH) -namespace=default -config=./cfg.yaml

echo Cleaning up...
k3d delete --name $K3D_NAME
rm ./kubernetes-diff-logger

echo All done.
