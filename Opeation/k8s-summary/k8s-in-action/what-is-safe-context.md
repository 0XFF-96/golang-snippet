### 什么是 安全 Context ?

1、为 Kubernetes 上运行的进程设定内核功能则需要于 Pod 内容器上的安全上下文中嵌套 capabilities 字段实现，

添加和移除内核能力还需要分别再向下一级嵌套使用 add 或 drop 字段
