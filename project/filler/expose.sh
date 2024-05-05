microk8s.kubectl delete services filler-0
sleep 1
microk8s.kubectl expose po filler-0  --type LoadBalancer --port 5001 --load-balancer-ip=192.168.169.58
microk8s.kubectl expose po filler-1  --type LoadBalancer --port 5001 --load-balancer-ip=192.168.169.59
microk8s.kubectl expose po filler-2  --type LoadBalancer --port 5001 --load-balancer-ip=192.168.169.60
microk8s kubectl get services 
