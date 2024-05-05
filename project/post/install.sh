helm delete post
helm install post bitnami/postgresql

sleep 10
./expose.sh
