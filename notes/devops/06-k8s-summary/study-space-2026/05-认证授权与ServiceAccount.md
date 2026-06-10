# 05 · 认证、授权与 ServiceAccount

> 主线:**你/你的程序凭什么能操作 K8s?**

## 三道关:认证 → 授权 → 准入
任何请求打到 API Server 都依次过三关,任一失败即拒:
```
请求 → [1.认证 Authentication] → [2.授权 Authorization] → [3.准入 Admission] → 写入 etcd
       你是谁?                    你能不能做?              合规校验/改写
```

## 一、两种身份:User Account vs ServiceAccount
| | User Account(用户账号) | ServiceAccount(服务账号) |
|---|---|---|
| 主体 | **人** | **程序 / Pod** |
| 来源 | K8s **不存储用户**,由外部管理(证书/OIDC/云厂商) | K8s 里的**标准资源对象**,存 etcd |
| 凭证 | kubeconfig 里的证书/token | 自动挂载到 Pod 的 token |
| 作用域 | 集群级 | **绑定到某 namespace** |

> 🔑 K8s 里没有"用户"这种资源对象(`kubectl get users` 查不到)。用户由认证凭证(证书 CN、OIDC sub)动态体现;而 SA 是实打实的对象。

## 二、认证:为什么你能用 kubectl?
答案在 **kubeconfig**(默认 `~/.kube/config`):
```yaml
clusters:   # 连哪个集群(API Server 地址 + 集群 CA)
users:      # ★ 你是谁(客户端证书:CN=用户名,O=用户组)
contexts:   # 把"集群+用户+namespace"绑成上下文,可切换
```
真相:你的 kubeconfig 有一份被集群 CA 信任的**客户端证书**,API Server 验证后读出 CN(用户名)和 O(组),这就是你的身份。
- 云厂商(阿里云 ACK/EKS)的 kubeconfig = 它帮你签发了一份带凭证的 kubeconfig,你写进 `~/.kube/config` 即可连云上集群。
- 认证方式:客户端证书(X.509)、Bearer Token(SA token)、OIDC(企业 SSO)、Webhook、云厂商插件。
```bash
kubectl config view        # 看当前配置
```

## 三、ServiceAccount:Pod 凭什么访问 API
- 给**程序**用的身份。⚠️ **不支持 spec 字段**,只指定名字,K8s 自动建关联凭证。
- token 注入 Pod:`/var/run/secrets/kubernetes.io/serviceaccount/`(含 ca.crt、namespace、token)。
- Pod 内程序用 token 当 Bearer Token 敲 API:
```bash
TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
curl --cacert .../ca.crt -H "Authorization: Bearer $TOKEN" \
  https://kubernetes.default.svc/api/v1/namespaces/default/pods
```
- 📌 **新版变化**:1.24+ 起 SA 不再自动建长期 Secret token,改用**绑定的、有时效、自动轮换**的投影 token(TokenRequest API),更安全。
- 三个自动化组件:ServiceAccount 准入控制器、Token 控制器、ServiceAccount 控制器。
- 给 Pod 指定 SA:
```yaml
spec:
  serviceAccountName: my-sa
  automountServiceAccountToken: false   # 不需访问 API 就关掉
```
- imagePullSecrets 挂到 SA → 该 SA 下所有 Pod 自动带拉私有镜像凭证。

## 四、授权:你能不能做?
内建插件:**RBAC**(生产唯一推荐)、Node(给 kubelet)、ABAC(过时)、Webhook。

### RBAC 四个核心对象
```
[Role / ClusterRole]  ──被──  [RoleBinding / ClusterRoleBinding]  ──绑给──  [User/Group/SA]
   定义"能做什么"                  把权限授予某身份
```
- **Role**:namespace 级权限集;**ClusterRole**:集群级。
- **RoleBinding**:在某 namespace 内授予;**ClusterRoleBinding**:集群范围授予。

示例:让 my-sa 只能读 default 空间的 Pod
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata: { namespace: default, name: pod-reader }
rules:
  - apiGroups: [""]
    resources: ["pods"]
    verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata: { namespace: default, name: read-pods }
subjects:
  - { kind: ServiceAccount, name: my-sa, namespace: default }
roleRef:
  { kind: Role, name: pod-reader, apiGroup: rbac.authorization.k8s.io }
```

### RBAC 生产铁律
- 🔴 **最小权限**:绝不给程序 `cluster-admin`(被攻破 = 整个集群沦陷)。
- 默认 `default` SA 应零权限;需要权限的程序单独建 SA + 精确授权。
- `verbs` 收紧:能 `get/list` 就别给 `*`。
- 自检:
```bash
kubectl auth can-i create pods
kubectl auth can-i list secrets --as=system:serviceaccount:default:my-sa
```

## 一句话总结
> kubectl 能操作集群 = kubeconfig 里有被集群 CA 信任的证书(认证)+ 该身份在 RBAC 被授权(授权);
> Pod 内程序能访问 API = 被注入了 SA 的 token(认证)+ 该 SA 有 RBAC 授权;
> 拉私有镜像 = imagePullSecrets 挂到 SA 上。
