microk8s kubectl delete -f .
microk8s kubectl create -f .
sleep 10
microk8s kubectl wait pods mysql-0  --for condition=Ready --timeout=90s
microk8s kubectl wait pods mysql-1  --for condition=Ready --timeout=90s
microk8s kubectl wait pods mysql-2  --for condition=Ready --timeout=90s
./expose.sh
microk8s kubectl delete po mysql-client
./de_db.sh
./create_db.sh
