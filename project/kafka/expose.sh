#microk8s.kubectl delete services web
#sleep 1
#microk8s.kubectl delete svc  kafka-0  
microk8s.kubectl delete svc zookeeper-0  zookeeper-1  zookeeper-2 
#microk8s.kubectl expose po kafka-0  --type LoadBalancer --port 9092 
microk8s.kubectl expose po zookeeper-0  --type LoadBalancer --port 2181 --load-balancer-ip=192.168.169.60
microk8s.kubectl expose po zookeeper-1  --type LoadBalancer --port 2181 --load-balancer-ip=192.168.169.61
microk8s.kubectl expose po zookeeper-2  --type LoadBalancer --port 2181 --load-balancer-ip=192.168.169.62
microk8s kubectl get services 
