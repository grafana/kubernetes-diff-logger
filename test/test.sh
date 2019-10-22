K3D_NAME=k8s-diff-logger

# requires k3d
k3d create --name $K3D_NAME

go build ../

sleep 5

./kubernetes-diff-logger -kubeconfig=$(k3d get-kubeconfig --name="${K3D_NAME}") -namespace=default

k3d delete --name $K3D_NAME
rm ./kubernetes-diff-logger