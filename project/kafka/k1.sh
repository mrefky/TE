#helm install kafka bitnami/kafka --set zookeeper.enabled=false --set replicaCount=3 --set externalZookeeper.servers=zookeeper.default.svc.cluster.local 
#./Nclient.sh
helm install kafka bitnami/kafka \
      --set zookeeper.enabled=false \
        --set replicaCount=3 \
          --set externalZookeeper.servers=zookeeper.default.svc.cluster.local \
 --set  externalAccess.enabled=true \
 --set externalAccess.service.type=LoadBalancer \
 --set deleteTopicEnable=true \
 --set externalAccess.service.port=19092 \
 --set externalAccess.service.loadBalancerIPs[0]='192.168.1.70' \
 --set externalAccess.service.loadBalancerIPs[1]='192.168.1.71' \
 --set externalAccess.service.loadBalancerIPs[2]='192.168.1.72'\
 --set auth.clientProtocol=sasl\
 --set auth.saslMechanisms=PLAIN\
 --set authorizer.class.name=kafka.security.auth.SimpleAclAuthorizer


