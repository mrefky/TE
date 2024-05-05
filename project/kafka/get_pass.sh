microk8s kubectl get secret kafka-jaas --namespace default -o jsonpath='{.data.client-passwords}' | base64 -d | cut -d , -f 1
