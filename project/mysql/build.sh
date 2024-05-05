echo "-------------------------------------"
echo "---   Delete mysql ------------------"
echo "-------------------------------------" 
microk8s kubectl delete -f .
microk8s kubectl create -f .
sleep 10
echo "-------------------------------------"
echo "---   Create mysql statefullset -----"
echo "-------------------------------------" 

microk8s kubectl wait pods mysql-0  --for condition=Ready --timeout=90s
microk8s kubectl wait pods mysql-1  --for condition=Ready --timeout=90s
microk8s kubectl wait pods mysql-2  --for condition=Ready --timeout=90s
./expose.sh
microk8s kubectl delete po mysql-client
echo "-------------------------------------"
echo "---   Delete database  --------------"
echo "-------------------------------------" 

./de_db.sh

echo "-------------------------------------"
echo "---   Create DB    ------------------"
echo "-------------------------------------" 

./create_db.sh
