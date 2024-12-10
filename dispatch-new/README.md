#  Install rabbitmq cluster

```
kubectl apply -f "https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml"
```

## Deploy a rabbitmq cluster
```yaml
apiVersion: rabbitmq.com/v1beta1
kind: RabbitmqCluster
metadata:
  name: my-rabbitmq-cluster
spec:
  replicas: 1  # Adjust as needed for more replicas
  service:
    type: NodePort  # or NodePort, depending on your access needs
```

## install rabbit cluster
```shell
ks create -f rabbit-cluster.yaml
kubectl get pods -l app.kubernetes.io/name=my-rabbitmq-cluster
```

## create a guest user as administrator
```shell
ks exec -it       -- bash   
rabbitmqctl add_user guest guest 
rabbitmqctl set_permissions -p / guest ".*" ".*" ".*"
rabbitmqctl set_user_tags guest administrator
# if you wamt change the password
rabbitmqctl change_password guest 12345678
# delete user
rabbitmqctl delete_user guest
# list users
rabbitmqctl list_users
# checking permissions
rabbitmqctl list_user_permissions guest
```
## Console
```shell
kubectl port-forward svc/my-rabbitmq-cluster 15672:15672 --address=0.0.0.0
# console web  http://<ip>:15672
```