microk8s kubectl delete po kafka-client 
microk8s kubectl run kafka-client --restart='Never' --image docker.io/bitnami/kafka:3.3.2-debian-11-r0 --namespace default --command -- sleep infinity
