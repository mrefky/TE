microk8s kubectl delete deploy --all

./clean.sh
sudo rm /mylogs/*
#cd ~/project/kafka
#./delete_zookeeper.sh
#./delete_kafka.sh
#./install_zookeeper.sh
#./install_kafka.sh
#cd ~/project/simple_kaka
#./expose.sh
#./clean.sh
cd ~/project/sql/
./build.sh
cd ~/project/matcher/
./build.sh
