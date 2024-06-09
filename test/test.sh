#!/usr/bin/env bash

K3D_NAME=k8s-diff-logger
CONFIG="k3d kubeconfig get $K3D_NAME"
CONFIG_PATH="kubeconfig.yaml"

function cleanup {
  echo Cleaning up...
  k3d cluster delete $K3D_NAME

  if [[ -f ./kubernetes-diff-logger ]]; then
    rm ./kubernetes-diff-logger
  fi

  if [[ -f $CONFIG_PATH ]]; then
    rm $CONFIG_PATH
  fi

  trap - INT TERM
}

trap cleanup INT TERM

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
eval $CONFIG > $CONFIG_PATH
echo Config is available at $CONFIG_PATH

echo
echo Running kubernetes-diff-logger. Make changes to the cluster in another shell. Press Ctrl-C when finished.
./kubernetes-diff-logger -kubeconfig=$CONFIG_PATH -namespace=default -config=./cfg.yaml -log-added=true -log-deleted=true

echo All done.
