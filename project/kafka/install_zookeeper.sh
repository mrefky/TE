helm repo add bitnami https://charts.bitnami.com/bitnami
#helm install zookeeper bitnami/zookeeper --set replicaCount=3 --set auth.enabled=false --set allowAnonymousLogin=true
helm install zookeeper bitnami/zookeeper \
	  --set replicaCount=3 \
 --set persistence.size=30Gi \
 --set persistence.dataLogDir.size=30Gi
 #--set auth.enabled=false\
#--set allowAnonymousLogin=true
