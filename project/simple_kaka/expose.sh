#microk8s.kubectl delete services --all
#sleep 1

microk8s.kubectl delete -f  kafka.all.2.yaml 
microk8s.kubectl create -f  kafka.all.2.yaml

microk8s.kubectl expose po zookeeper-0 --type LoadBalancer --port 2181
microk8s.kubectl expose po kafka-0 --type LoadBalancer --port 29092
microk8s kubectl get services
