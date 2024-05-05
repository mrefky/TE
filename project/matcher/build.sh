microk8s.kubectl delete -f matcher2.yaml
sudo rm /mylogs/matcher.log
sleep 3
go build .
docker build -t localhost:32000/slim .
docker push localhost:32000/slim
sleep 1
microk8s.kubectl create -f matcher2.yaml
