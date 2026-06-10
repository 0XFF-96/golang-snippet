### Deployment


1、https://kubernetes.io/docs/concepts/workloads/controllers/deployment/

Do not overlap labels or selectors with other controllers. ! k8s don’t stop you from overlapping, and if multiple controllers have overlapping selectors those controllers might conflict and behave unexpectedly .

2、by default 75% 和 25 % 的 pod 是会超过 deployment 设置的阈值。 RollingUpdateStrategy:  25% max unavailable, 25% max surge
3、为什么 update deployment 的 spec 之后，出现两个 rs ?
4、In API version apps/v1, a Deployment's label selector is immutable after it gets created
5、revision is created when the deployment’s pod template (.spec.template) is changed . 
6、kubectl autoscale deployment.v1.apps/nginx-deployment --min=10 --max=15 --cpu-percent=80
HPA: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/ 。 
7、Pausing and Resuming, 为什么需要 pausing ?
8、太多了….没有看完。 
