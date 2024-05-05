sudo rm /mylogs/sql.log
sleep 10 
microk8s kubectl delete -f sql2-1.yaml
go build .
docker build -t localhost:32000/sql .
docker push localhost:32000/sql
#cp ./sql ~/TE/.
#cd ~/TE
#./reset.sh
sleep 1
microk8s kubectl create -f sql2-1.yaml
