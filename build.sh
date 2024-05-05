
chmod a+x *sh
chmod a+x */*sh

microk8s kubectl delete sts mysql
microk8s kubectl delete sts kafka
microk8s kubectl delete sts zookeeper
microk8s kubectl delete sts --all
microk8s kubectl delete deploy --all
microk8s kubectl delete svc --all
microk8s kubectl delete pvc --all

sleep 1
cd ~/TE/project/mysql
./build.sh
cd ~/TE/project/kafka
./build.sh
cd ~/TE/project/control
./resetall.sh
cd ~/TE/project/filler
./build.sh 
cd ~/TE/project/mysql-view
./build.sh
k9s
