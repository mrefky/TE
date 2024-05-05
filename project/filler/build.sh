microk8s kubectl delete -f filler.yaml
go build .
./create_filler.sh
docker build -t localhost:32000/filler .
docker push localhost:32000/filler

sleep 1
microk8s kubectl create -f filler.yaml
microk8s kubectl wait pods filler-0  --for condition=Ready --timeout=90s
./expose.sh
