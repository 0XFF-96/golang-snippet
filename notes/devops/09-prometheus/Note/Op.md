# Promethues

## Outline 
- 查询语句文档 https://prometheus.io/docs/prometheus/latest/querying/basics/ 
- 

1、https://xuchao918.github.io/2020/02/14/Kubernetes%E7%9B%91%E6%8E%A7%E5%91%8A%E8%AD%A6%EF%BC%88%E4%B8%80%EF%BC%89/?utm_source=tuicool&utm_medium=referral

 

### 资源

1、https://songjiayang.gitbooks.io/prometheus/content/concepts/data-model.html
2、https://hulining.gitbook.io/prometheus/introduction/media
3、https://erdong.site/prometheus-notes/chapter01-Basic/1.5-data-model.html

### k8s and prometheus 

1、https://github.com/kubernetes/kube-state-metrics/tree/master/Documentation
2、https://github.com/kubernetes/kube-state-metricsPrometheus 支持通过多个监控目标采集 Kubernetes, 
3、https://prometheus.io/docs/practices/instrumentation/ 


### case 

1、统计一下发送了多少次请求。
2、统计某个接口的某个分支条件被执行了多少次。 例如，上班 【抖音】广告平台的次留接口执行次数。【能否查出来过去 24 小时被执行了多少次？】
【什么时候 counter 数据会被重置为 0 值】 

### Data Model 数据模型

1、https://prometheus.io/docs/concepts/data_model/
2、Metric Names And Labels 
1、Every time series is uniquely identified by its metric name and optional key-value pairs called labels
2、labels: Labels enable Prometheus's dimensional data model 
3、Changing any label value, will create a new time series. 
4、<metric name>{<label name>=<label value>, ...}5、Label:   Use labels to differentiate the characteristics of the thing that is being measured。 

Job and instances: 1、Reference
1、time_series: [WikiPedia article about Time series](https://en.wikipedia.org/wiki/Time_series) 。 
2、https://prometheus.io/docs/practices/naming/
3、https://prometheus.io/docs/concepts/jobs_instances/

### OverView 

Outline:
1、What is prometheus ? features、metrics、components 
2、What does it fit or not fit ?

基本概念：
1、Prometheus collects and stores its metrics as time series data, i.e. metrics information is stored with the timestamp at which it was recorded, alongside optional key-value pairs called labels.

2、Prometheus works well for recording any purely numeric time series
3、trade-off : Prometheus offers a richer data model and query language, in addition to being easier to run and integrate into your environment. If you want a clustered solution that can hold historical data long term, Graphite may be a better choice.4、Comparsion to alternatives. Graphite、InfluxDB、

5、In layperson（Insiders） terms, Resources:
1、https://prometheus.io/docs/introduction/faq/   ?
2、how to troubleshooting ?
3、talks and research , https://prometheus.io/docs/introduction/media/ 。 


### 部署 Op 

1、https://prometheus.io/docs/prometheus/latest/getting_started/
2、端口占用情况， https://zhuanlan.zhihu.com/p/355376119#:~:text=%E4%BB%8A%E5%A4%A9%E5%B0%8F%E7%BC%96%E7%BB%99%E5%A4%A7%E5%AE%B6,kill%20pid%EF%BC%89%E8%AE%B0%E5%BE%97%E6%94%B6%E8%97%8F%E5%93%A6%EF%BC%81 
3、从 docker 启动一个镜像。 docker run --name prometheus -d -p 127.0.0.1:9091:9090 prom/prometheus 。 
4、These labels designate different latency percentiles and target group intervals,
5、从二进制部署时，踩了一个配置文件的坑。现在还没有找到 listen-address 配置文件的具体方式。 
6、学会配置如何 scrape 一个组件。 学会配置如何合并 time series metric  ?7、[www.cnblogs.com/fsckzy/p/13335173.html](https://www.cnblogs.com/fsckzy/p/13335173.html)

疑惑🤔：
1、为什么不能够热更新？需要每次配置 promethueus 的配置然后 xxxx. 

### 资源

1、hello world 的 prometheus 
2、这位老哥看过的 prometheus 资源比较多， [www.bboy.app/page/2/](https://www.bboy.app/page/2/) 。 


