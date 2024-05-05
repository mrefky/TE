./delete_zookeeper.sh
./delete_kafka.sh
microk8s.kubectl delete sts --all
microk8s.kubectl delete pvc --all
#./install_zookeeper.sh
./install_kafka.sh
./expose.sh
