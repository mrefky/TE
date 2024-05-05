#microk8s.kubectl delete svc mysql-0
#microk8s.kubectl delete svc mysql-1
#microk8s.kubectl delete svc mysql-2

microk8s.kubectl expose po post-postgresql-0   --type LoadBalancer --port 3306 --load-balancer-ip=192.168.1.50
#microk8s.kubectl expose po mysql-1  --type LoadBalancer --port 3306 --load-balancer-ip=192.168.1.51
#microk8s.kubectl expose po mysql-2  --type LoadBalancer --port 3306 --load-balancer-ip=192.168.1.52
microk8s kubectl get services
