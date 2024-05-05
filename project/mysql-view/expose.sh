microk8s.kubectl delete services web
sleep 1
microk8s.kubectl expose po web  --type LoadBalancer --port 8080 --load-balancer-ip=192.168.169.90
microk8s kubectl get services 
