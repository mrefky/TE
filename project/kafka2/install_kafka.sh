helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
microk8s.kubectl create ns kafka
helm install kafka --values values.yaml -n kafka bitnami/kafka
