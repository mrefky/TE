# TE
# System Overview

A simplified description is displayed in the below picture where the client application sends orders to a Kafka cluster either directly or through a provided fix gateway service we named it as “Filler”; Kafka stores these orders in a topic named “orders”. Then the matching engine received these orders; match them and send the resultant trades back to Kafka which stores it in a topic named “trades”.
Another service named “SQL-Updater” collects both trades and orders and send them to a MYSQL database cluster.
All these services are deployed into Kubernetes cloud (Microk8s).
The connection between the DB and matching engine are required to ensure that the matching engine will only accept orders for a registered security and for getting all market data specifications and also to be able to recover if it fails by reading and rebuilding the latest order books for all securities (will be explained in details later.



![arch](./arch.jpg?raw=true "Arch")

# Testing
* stress testing

    cd ~/TE/project/stress3
  
    go run .
    
* predefined set of orders
* Using FIx

# Generated STS
![arch](./sts.jpg?raw=true "Arch")

# Generated pods

![arch](./pods.jpg?raw=true "Arch")

# Generated Services

![arch](./svc.jpg?raw=true "Arch")
