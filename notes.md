
## Create a Cluster

#### GKE

- GKE is the easiest to get running
  - assuming you have the gcloud tools installed and a project setup

#### Docker Edge

- Docker edge is fairly easy to install
- You have to enable the kubernetes cluster seperately
- Most things work
- Comes with a nice GUI

#### minikube

- Harder to install
- LoadBalancer Services do not expose a external IP

## Create Clusters Demo

```
gcloud container clusters create k8s-basics-talk-cluster --zone=us-central1-a --cluster-version 1.10
kubectl config get-contexts

minikube start
kubectl config get-contexts

# show docker edge GUI for creating cluster

kubectl config use-context minikube
```

## What is a Cluster

- Detail was removed from this diagram for clairity

#### Masters

- masters orchestrate the cluster
- etcd is the store
- Most things only talk to the apiserver
- The scheduler talks to the api and makes scheduling decisions based on
  available nodes, capacity, affinity, etc.

#### Nodes

- Nodes run the workloads
- The kubelet uses a container runtime, typically docker.
- kube-proxy exists on the nodes to forward traffic to the right node that has
  the workload.

## Resource Types

- Pod: Unit of scheduling
- RelicaSets: Create multiple pods of the same app
- Deployments: Allow for rolling deploys of an app
- Services: Load balance traffic onto Pods
- Namespace: Isolate groups of apps

## Nginx Demo

```
kubectl run nginx --image nginx
kubectl get pods
kubectl scale deployment nginx --replicas 5
kubectl get pods
kubectl expose deployment nginx --port 80 --type LoadBalancer
kubectl get services

kubectl get service nginx -o json | jq .status.loadBalancer.ingress[].ip -r
ip=$(!!)
curl $ip

kubectl delete service nginx
kubectl delete deployment nginx
```

## Custom App Demo

```
# describe the app source
# briefly describe the Dockerfile

# get a watch going
watch -n .1 kubectl get all

cat pod.yml
kubectl apply -f pod.yml

cat service.yml
kubectl apply -f service.yml
kubectl get service app -o json | jq .status.loadBalancer.ingress[].ip -r
ip=$(!!)
curl $ip
curl $ip/crash
curl $ip

# get a watch going
watch -n .1 kubectl logs app

curl $ip/crash

kubectl exec -it app /bin/sh
ps aux
netstat -ln


cat replicaset.yml
kubectl apply -f replicaset.yml
kubectl delete pod app
kubectl delete rs app


cat deployment.yml
kubectl apply -f deployment.yml

# get a watch going
kubectl get service app -o json | jq .status.loadBalancer.ingress[].ip -r
watch -n .1 curl -s $(!!)

kubectl edit deployment app
# change image to nginx
```
