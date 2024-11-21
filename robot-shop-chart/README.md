# install 
```shell
helm package ./robot-shop-chart/
curl -u admin:xxxxxx http://170.75.162.227:32510/repository/robotshop-helm/ --upload-file robot-shop-chart-0.1.0.tgz

helm repo add robotshop-helm  http://170.75.162.227:32510/repository/robotshop-helm/ --username admin --password xxxxxx

helm install v.0.1.0  robotshop-helm/robot-shop-chart

kubectl create secret docker-registry nexus-cred \
--docker-server=nexus:30999 \
--docker-username=admin \
--docker-password=xxxxx \
--docker-email=dockerlite@gmail.com
```

## refresh helm package
```shell
cd /home/ubuntu/robot-shop/leanpub
kind delete cluster --name robotshop
kind create cluster --name robotshop --config kind-config-cluster.yml 
cd ..
helm package robot-shop-chart
curl -u admin:xxxxxx http://170.75.162.227:32510/repository/robotshop-helm/ --upload-file robot-shop-chart-0.1.0.tgz
helm install v.0.1.0  robotshop-helm/robot-shop-chart
ks get pod
watch 'kubectl get pods'
    
```


