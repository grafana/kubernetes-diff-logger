#!/usr/bin/env bash
K3D_NAME=k8s-diff-logger
CONFIG="k3d kubeconfig get $K3D_NAME"
CONFIG_PATH="kubeconfig.yaml"

# requires k3d
k3d cluster create $K3D_NAME --api-port=6444

echo Building executable... 
go build ../

echo -n "Ensuring k3d is running..."
while true; do
  k3d cluster list 2>&1 | grep ".*$K3D_NAME.*1/1" >/dev/null && echo "done" && break \
    || (echo -n . && sleep 1)
done

echo -n "Getting kubeconfig..."
while true; do
  eval $CONFIG 2>&1 | grep "k3d-$K3D_NAME" >/dev/null && echo done && break \
    || (echo -n . && sleep 1)
done
echo $(eval $CONFIG) > $CONFIG_PATH
echo Config is available at $CONFIG_PATH

echo Running kubernetes-diff-logger...
./kubernetes-diff-logger -kubeconfig=$CONFIG_PATH -namespace=default -config=./cfg.yaml

echo Cleaning up...
k3d cluster delete $K3D_NAME
rm ./kubernetes-diff-logger
rm $CONFIG_PATH

echo All done.
