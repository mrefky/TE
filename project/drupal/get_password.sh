 echo Password: $(microk8s.kubectl get secret --namespace default my-release-drupal -o jsonpath="{.data.drupal-password}" | base64 -d)
