microk8s kubectl delete po mysql-client-loop

microk8s  kubectl run mysql-client-loop --image=mysql:5.7 -i -t --rm --restart=Never --  bash -ic "while sleep 1; do mysql -h mysql-read -e 'SELECT (SELECT @@server_id )as Server ,(     SELECT COUNT(*)     FROM   test.orders ) AS orders, ( select count(*) from test.trades) As trades'; done"
