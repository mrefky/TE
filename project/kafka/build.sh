./delete_kafka.sh
microk8s.kubectl wait pod,svc,cm,pvc  --all --for=condition=Ready  --all-namespaces
./install_kafka.sh
microk8s.kubectl wait pod,svc,cm,pvc  --all --for=condition=Ready  --all-namespaces
./expose.sh
