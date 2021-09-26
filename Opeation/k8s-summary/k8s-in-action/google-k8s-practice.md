### google 云平台

1. 登录地址 https://console.cloud.google.com/billing/016736-8DD946-CA6F84 
2. 教程： http://cloud.google.com/kubernetes-engine/docs/tutorials/

### google 部署微服务

一、拉代码
```
git clone \
    https://github.com/GoogleCloudPlatform/microservices-demo.git
```


cd microservices-demo
watch kubectl get \
    pods,deployments,services


kubectl describe pod -l app=frontend

kubectl describe pod -l app=frontend

CPU='"650m"'; kubectl get \
    deployment frontend -o json | \
    jq \
    ".spec.template.spec.containers[0].resources.requests.cpu \
    = $CPU" | kubectl apply -f -


kubectl scale deployment frontend \
    --replicas=3

kubectl scale deployment frontend \
    --replicas=1kubectl autoscale deployment \
    frontend --cpu-percent=5 \
    --min=1 --max=3kubectl autoscale deployment \
    frontend --cpu-percent=5 \
    --min=1 --max=3kubectl scale deployment \
    loadgenerator --replicas=5watch kubectl get pods,hpa

REGION=$(gcloud container clusters \
    list | grep \
    my-autopilot-cluster | awk \
    '{print $2}')
