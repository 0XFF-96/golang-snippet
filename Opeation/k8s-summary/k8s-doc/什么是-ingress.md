### 什么是 ingress ?

1、kubectl get ingress
2、kubectl describe ingress dc-staging-sundayf
3、前端应该有一个外部的负载均衡设备，接收并调度此类请求。  4、https://github.com/kubernetes/ingress-nginx 。  5、Ingress. 是 service 的反向代理。 （精辟） （从哪里来的？）

annotations: 是什么意思？ 
    # Enable canary and send traffic with headder x-app-canary to version 2
    http://nginx.ingress.kubernetes.io/canary: "true"
    http://nginx.ingress.kubernetes.io/canary-by-header: "x-app-canary"
    http://nginx.ingress.kubernetes.io/canary-by-header-value: "true"
    
    