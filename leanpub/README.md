# README

## install Docker on ubuntu 24.04
```shell
sudo apt update
sudo apt install -y curl apt-transport-https ca-certificates software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install docker-ce -y
sudo systemctl status docker
sudo usermod -aG docker $USER
# log out log in again
docker ps # check
```

## install docker-compose
```shell
sudo curl -L "https://github.com/docker/compose/releases/download/v2.29.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/bin/docker-compose
sudo chmod +x /usr/bin/docker-compose 
docker-compose version 
```

## install robot-shop 
```shell
git clone https://github.com/crunchy-devops/robot-shop.git
cd robot-shop
docker-compose build
docker-compose up -d
```

## install on kubernetes kind
```shell
wget https://github.com/kubernetes-sigs/kind/releases/download/v0.24.0/kind-linux-amd64
mv kind-linux-amd64 kind
chmod +x kind
sudo mv kind /usr/local/bin/kind
kind version # should be  version 0.24.0
```
## install kubectl
```shell
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/kubectl
alias ks='kubectl'
source <(kubectl completion bash | sed s/kubectl/ks/g)
```

## Create a cluster
```shell
kind create cluster --name awx --config kind-config-cluster.yml
ks version # should be version  v1.31.1+
ks get nodes # see one control-plane and 3 workers
```

## install AWX
```shell
cd
git clone https://github.com/ansible/awx-operator.git
cd awx-operator/
git checkout tags/2.19.1
git log --oneline  # HEAD should be on tag 2.19.1 #hash dd37ebd
q
```

## Troubleshooting to prevent job template failure in AWX
```shell
echo fs.inotify.max_user_watches=655360 | sudo tee -a /etc/sysctl.conf
echo fs.inotify.max_user_instances=1280 | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```

### Manually create file kustomization.yaml
```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  # Find the latest tag here: https://github.com/ansible/awx-operator/releases
  - github.com/ansible/awx-operator/config/default?ref=2.19.1
  - awx-demo.yml
# Set the image tags to match the git version from above
images:
  - name: quay.io/ansible/awx-operator
    newTag: 2.19.1

# Specify a custom namespace in which to install AWX
namespace: awx
```
```ks create ns awx```
```ks apply -k . ```  run twice this command

Wait 15 minutes
And check with ```ks get pod -A``` # all K8s objects should be running, completed

## User AWX
username admin
Password uses the commande below
```shell
kubectl get secret -n awx  awx-demo-admin-password -o jsonpath="{.data.password}" | base64 --decode ; echo
```
## Web access
```
kubectl port-forward -n awx service/awx-demo-service 30880:80 --address='0.0.0.0'
```
access to AWX with http://<ip>:30880



## Check images size stats
```shell
docker images --all --filter=reference='robotshop/*:ebde46fae5ae' --format "{{.Repository}}\t{{.Size}}" >/bitnami/jenkins/home/checkimages/build-11
```