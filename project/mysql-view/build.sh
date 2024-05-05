go build .
docker build -t localhost:32000/web .
docker push localhost:32000/web
microk8s.kubectl delete -f web.yaml
microk8s.kubectl create -f web.yaml
./expose.sh
#cp ./sql ~/TE/.
#cd ~/TE
#./reset.sh
