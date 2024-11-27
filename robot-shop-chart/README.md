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

## useful commands
```shell
echo -n 'admin' | base64
```

## install nfs in kubernetes kind
```shell
sudo apt update
sudo apt install -y nfs-kernel-server
sudo mkdir -p /mnt/nfs_share
sudo chown nobody:nogroup /mnt/nfs_share
vi /etc/exports
sudo vi /etc/exports
sudo exportfs -rav
sudo exportfs -v
sudo systemctl enable --now nfs-server
# change kind-config file


helm repo add nfs-subdir-external-provisioner https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/
helm repo update
helm install nfs-client nfs-subdir-external-provisioner/nfs-subdir-external-provisioner \
  --set nfs.server=172.16.0.8 \
  --set nfs.path=/mnt/dump

```

