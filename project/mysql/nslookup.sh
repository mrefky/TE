
 microk8s kubectl apply -f https://k8s.io/examples/admin/dns/dnsutils.yaml
 sleep 10
  microk8s kubectl exec -i -t dnsutils -- nslookup mysql-0 


