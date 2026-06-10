# K8s 学习空间(2026 整理版)

> 本目录是与 AI 结对梳理后的系统化笔记,把零散的提纲式笔记补全成可复习的完整知识。
> 整理日期:2026-06。

## 一句话心智模型

> **Kubernetes 是一个"调和引擎":你声明期望状态(spec),控制器不断让现实(status)向它对齐。**
> 声明工作负载 + 配置 + 存储 → 控制器调和 → 标签把一切串起来 → Service/Ingress 让它可达 → Namespace + RBAC 保证有序和安全。

## 目录索引

| # | 文件 | 内容 |
|---|------|------|
| 01 | [核心概念](01-核心概念.md) | 10 大基础概念、对象模型、spec/status、调和循环、标签、Namespace |
| 02 | [工作负载](02-工作负载.md) | Pod、ReplicaSet、Deployment、StatefulSet、DaemonSet、Job、CronJob |
| 03 | [服务发现与网络](03-服务发现与网络.md) | Service、Endpoints、kube-proxy、Ingress、三个 Port、DNS |
| 04 | [配置与密钥](04-配置与密钥.md) | ConfigMap、Secret、热更新、default-token |
| 05 | [认证授权与 ServiceAccount](05-认证授权与ServiceAccount.md) | 三道关、kubeconfig、RBAC、imagePullSecrets |
| 06 | [排错手册](06-排错手册.md) | 网络、DNS、服务连通性、端口占用排查 |
| 07 | [生产实践与 Argo](07-生产实践与Argo.md) | 资源/探针/优雅关闭/调度/伸缩/安全/可观测性、Argo 家族 |

## 资源选型速查

| 需求 | 资源 |
|------|------|
| 无状态 Web/API | Deployment |
| 数据库、有身份的集群应用 | StatefulSet + Headless Service |
| 每个节点的日志/监控 agent | DaemonSet |
| 一次性批处理 | Job |
| 定时任务 | CronJob |
| 集群内稳定访问一组 Pod | Service (ClusterIP) |
| 对集群外暴露 TCP 服务 | Service (LoadBalancer/NodePort) |
| 对集群外暴露 HTTP,按域名/路径路由 | Ingress |
