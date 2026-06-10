# 07 · 生产实践与 Argo

## 一、资源管理(第一大事故源)
```yaml
resources:
  requests: { cpu: 100m, memory: 256Mi }   # 用于调度,保证值
  limits:   { memory: 256Mi }              # 硬上限
```
- **内存 requests == limits**:可预测,避免节点 OOM 意外驱逐。内存不可压缩,超限即 `OOMKilled`。
- **慎设 CPU 硬 limit**:CPU 可压缩(被节流不被杀)。硬 limit 导致 CFS 节流,即使节点空闲也拖垮 p99。很多团队只设 CPU requests。
- **QoS 类**(驱逐顺序):`Guaranteed`(req==limit,最后被杀)> `Burstable` > `BestEffort`(最先杀)。关键负载跑 Guaranteed。

## 二、健康探针(三个都要对)
```yaml
livenessProbe:    # 失败则重启容器。保持"傻":只查进程是否卡死
  httpGet: { path: /healthz, port: 8080 }
  periodSeconds: 10
  failureThreshold: 3
readinessProbe:   # 失败则从 Service Endpoints 摘除。这是流量闸门
  httpGet: { path: /ready, port: 8080 }
startupProbe:     # 保护慢启动应用,避免 liveness 在启动期误杀
  httpGet: { path: /healthz, port: 8080 }
  failureThreshold: 30
  periodSeconds: 10
```
- ⚠️ **liveness 别查依赖**(如 DB):DB 抖动 → 全部 Pod 重启 → 级联故障。**liveness = 进程是否卡死,readiness = 能否服务流量**。
- 慢启动(JVM)没配 startupProbe → liveness 启动期误杀循环。

## 三、零停机部署与优雅关闭
```yaml
lifecycle:
  preStop: { exec: { command: ["sleep", "10"] } }   # 等 Endpoint 摘除传播完再死
terminationGracePeriodSeconds: 30
```
- 应用**必须处理 SIGTERM**:停收新连接、排空在途请求、再退出。
- 设 **PodDisruptionBudget** 防止节点排空/升级时一次性挂太多副本:
```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
spec:
  minAvailable: 2
  selector: { matchLabels: { app: myapp } }
```

## 四、调度与高可用
- **topologySpreadConstraints** 把副本分散到多区/多节点,别让所有副本落同一节点:
```yaml
topologySpreadConstraints:
  - maxSkew: 1
    topologyKey: topology.kubernetes.io/zone
    whenUnsatisfiable: ScheduleAnyway
    labelSelector: { matchLabels: { app: myapp } }
```
- 污点/容忍(taint/toleration)专用节点;节点亲和(nodeAffinity)钉实例类型。

## 五、自动伸缩
- **HPA**:别只按 CPU,按反映真实负载的指标(RPS/队列深度/自定义,经 KEDA)。`minReplicas ≥ 2`。
- **VPA**:推荐模式下用于 right-sizing;别和 HPA 用同一指标。
- **Cluster Autoscaler / Karpenter**:节点伸缩。AWS 上 Karpenter 更快、装箱更好、能整合节点。

## 六、安全基线
```yaml
securityContext:
  runAsNonRoot: true
  runAsUser: 10001
  readOnlyRootFilesystem: true
  allowPrivilegeEscalation: false
  capabilities: { drop: ["ALL"] }
  seccompProfile: { type: RuntimeDefault }
```
- **Pod Security Admission**(`restricted` profile)在 namespace 级。
- **NetworkPolicy** 默认拒绝再按需放行(否则任意 Pod 互通)。
- **RBAC** 最小权限;关闭不需要的 SA token 自动挂载。
- Secret 用 External Secrets/Vault,镜像钉 digest 不用 `:latest`,用 distroless,扫描镜像(Trivy)。

## 七、可观测性(四黄金信号)
延迟、流量、错误、饱和度。
- **Metrics**:Prometheus + Grafana,告警基于症状(错误率/延迟 SLO)而非原因(CPU 高)。
- **Logs**:结构化 JSON,经 Fluent Bit/Vector 发到 Loki/ELK。
- **Tracing**:OpenTelemetry。
- **首日必备告警**:CrashLoopBackOff、OOMKilled、Pending(不可调度)、PVC 将满、证书过期、节点 NotReady、HPA 触顶。

---

# Argo 家族

| 项目 | 用途 | 一句话 |
|------|------|--------|
| **Argo CD** | GitOps 持续部署 | Git 是唯一真相源,集群向 Git 对齐 |
| **Argo Rollouts** | 渐进式发布 | 金丝雀/蓝绿,替代原生 Deployment |
| **Argo Workflows** | 工作流引擎 | 跑 CI/CD、批处理、ML pipeline 的 DAG |
| **Argo Events** | 事件驱动 | 基于事件触发 Workflow |

## Argo CD —— GitOps 核心
- 声明式 + **拉取模式(Pull)**:Argo 跑在集群内主动拉 Git 对比,比外部 push `kubectl apply` 更安全。
- **持续对账**:默认每 3 分钟轮询,漂移就标 `OutOfSync`。
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata: { name: myapp-prod, namespace: argocd }
spec:
  project: prod
  source:
    repoURL: https://github.com/org/deploy.git
    targetRevision: main          # 生产建议钉 tag/commit,别用 HEAD
    path: overlays/prod
  destination:
    server: https://kubernetes.default.svc
    namespace: myapp
  syncPolicy:
    automated: { prune: true, selfHeal: true }
    syncOptions: [CreateNamespace=true, PruneLast=true]
    retry: { limit: 5, backoff: { duration: 5s, factor: 2, maxDuration: 3m } }
```
**生产关键点**:
- `selfHeal: true` 双刃剑:杜绝手动改集群,但救火时 `kubectl edit` 会被秒回滚 → 先停 auto-sync。
- **App of Apps / ApplicationSet**:批量管理多环境/多集群。
- `prune` 小心:误删 Git 文件会真删生产资源,重要资源加 `Prune=false`。
- **AppProject + RBAC** 做多租户隔离。
- **Sync Waves**(`argocd.argoproj.io/sync-wave`)控制部署顺序(先 DB 再 App)。
- **漂移检测**价值:有人半夜改集群立刻 `OutOfSync`,审计/合规关键。

## Argo Rollouts —— 安全发布
原生 Deployment 不能暂停/按比例放量/按指标自动回滚。用 `Rollout` 替代:
```yaml
strategy:
  canary:
    steps:
      - setWeight: 5
      - pause: { duration: 5m }
      - setWeight: 20
      - pause: {}              # 无限暂停,等人工确认
      - analysis: { templates: [{ templateName: success-rate }] }  # 不达标自动回滚
      - setWeight: 50
      - setWeight: 100
```
- **AnalysisTemplate** 是灵魂:对接 Prometheus/Datadog 自动判断错误率/p99,不达标自动回滚。
- 蓝绿(瞬时切流、回滚快、双倍资源) vs 金丝雀(逐步放量、省资源、较慢)。
- 需配合 Ingress/Service Mesh 做流量切分。

## Argo CD vs Flux
- 要 UI、团队协作、成熟多租户 → **Argo CD**。
- 追求极简、纯 GitOps 控制器 → **Flux**。

## 生产踩坑总结
1. `targetRevision` 别用 `HEAD`/`main`,钉具体 tag/commit。
2. Secret 别明文进 Git(Sealed Secrets / External Secrets / SOPS)。
3. Argo CD 自身 HA:repo-server、application-controller 多副本,大规模调 sharding。
4. 通知:`argocd-notifications` 把失败/Degraded 推 Slack/钉钉。
5. 回滚:UI 可 History and Rollback,但最干净的是 `git revert` 让 Git 当真相源。
