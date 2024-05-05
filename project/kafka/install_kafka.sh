#helm install kafka bitnami/kafka --set zookeeper.enabled=false --set replicaCount=3 --set externalZookeeper.servers=zookeeper.default.svc.cluster.local 
#./Nclient.sh
helm install kafka bitnami/kafka  --version="22.1.3" \
	  --set zookeeper.enabled=false \
	    --set replicaCount=3 \
 --set  externalAccess.enabled=true \
 --set externalAccess.service.type=LoadBalancer \
 --set deleteTopicEnable=true \
 --set persistence.size=20Gi \
 --set logPersistence.size=20Gi \
 --set externalAccess.service.port=19092 \
 --set externalAccess.service.loadBalancerIPs[0]='192.168.169.70' \
 --set externalAccess.service.loadBalancerIPs[1]='192.168.169.71' \
 --set externalAccess.service.loadBalancerIPs[2]='192.168.169.72'

