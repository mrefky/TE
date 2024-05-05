


microk8s kubectl delete sts mysql
microk8s kubectl delete sts kafka
microk8s kubectl delete sts zookeeper
microk8s kubectl delete sts --all
microk8s kubectl delete deploy --all
microk8s kubectl delete svc --all
microk8s kubectl delete pvc --all

sleep 1
cd /home/mrefky/project/mysql
./build.sh
cd /home/mrefky/project/kafka
./build.sh
cd /home/mrefky/project/control
./resetall.sh
cd /home/mrefky/project/filler
./build.sh 
cd /home/mrefky/project/mysql-view
./build.sh
#/home/mrefky/project/status.sh
