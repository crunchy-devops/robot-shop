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
# Added this package in your hosts
sudo apt update
sudo apt install -y nfs-kernel-server
sudo mkdir -p /mnt/nfs_share
sudo mkdir -p /mnt/dump
sudo chmod 777  /mnt/nfs_share
sudo chmod 777 /mnt/dump
sudo chown nobody:nogroup /mnt/nfs_share
sudo chown nobody:nogroup /mnt/dump

sudo vi /etc/exports
# Added the following lines
#/mnt/nfs_share *(rw,sync,no_subtree_check,no_root_squash)
#/mnt/dump  *(rw,sync,no_subtree_check,no_root_squash)
sudo exportfs -rav
sudo exportfs -v
sudo systemctl enable --now nfs-server
# change kind-config file
# Added on all nodes the extraMounts
- role: control-plane
  extraMounts:
    - hostPath: /mnt/nfs_share
      containerPath: /mnt/nfs_share
    - hostPath: /mnt/dump
      containerPath: /mnt/dump
# first solution only for one nfs share directory
helm repo add nfs-subdir-external-provisioner https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/
helm repo update
helm install nfs-client nfs-subdir-external-provisioner/nfs-subdir-external-provisioner \
  --set nfs.server=172.16.0.8 \
  --set nfs.path=/mnt/dump
# Second solution you can have multi shared directories
curl -skSL https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/v4.9.0/deploy/install-driver.sh | bash -s v4.9.0 --
# usage of storage class
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: nfs-csi-server1
provisioner: nfs.csi.k8s.io
parameters:
  server: 172.16.0.8
  share: /mnt/nfs_share
  subDir: "data"
reclaimPolicy: Delete
volumeBindingMode: Immediate
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nfs-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 1Gi
  storageClassName: nfs-csi-server1

```