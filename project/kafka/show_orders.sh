microk8s kubectl exec --tty -i kafka-client --namespace default -- kafka-console-consumer.sh \
            --bootstrap-server kafka.default.svc.cluster.local:9092 \
            --topic orders \
            --from-beginning
