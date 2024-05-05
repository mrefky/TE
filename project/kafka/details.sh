microk8s.kubectl get svc --namespace default -l "app.kubernetes.io/name=kafka,app.kubernetes.io/instance=kafka,app.kubernetes.io/component=kafka,pod" -w
