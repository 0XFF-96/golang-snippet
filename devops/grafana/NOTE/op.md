
### 什么是 Grafana 的 Data Link 

1、如何使用 DataLink ?

### 利用 rsync 同步 Grafana 的数据

rsync -azvhP root@xxxcom:/var/lib/grafana/grafana.db ~/Downloads/grafana.db 。

### 什么是 grafana 的 transformation

1、https://grafana.com/docs/grafana/latest/panels/transformations/
2、

### Grafana 的部署和运维方式

1、是通过 ansible 的 galaxy 的方式进行的
2、grafana 的 storage 使用了 sqlite3 数据库。 
3、部署在 daycam-001 机器上。 
4、daycamsqlite3 ？
5、如何查看 grafana server 的相关日志。
6、需要更改我们线上的 grafana 的账号。 
7、通过 make ssh config 的方式，怎么去 revoke 一个人的 ssh 访问权限？

### Grafana 技巧

1、[www.qikqiak.com/post/use-crd-create-grafana-dashboard/](https://www.qikqiak.com/post/use-crd-create-grafana-dashboard/) ， 利用 k8s crd 创建 grafana dashboard 。 
2、从云端已经做好的 Dashboard , 引入。 Importing a dashboard 。 https://grafana.com/docs/grafana/v7.5/dashboards/export-import/ 。 
https://grafana.com/grafana/dashboards ， 在这个 Dashboard 里面。 
4、https://github.com/grafana/grafana/blob/main/docs/sources/best-practices/best-practices-for-creating-dashboards.md 。 
5、我们的 grafana 是配置在哪里的？如何被部署的？

### Grafana 的模板变量 

1、如何更快修改 query 变量
2、如何直接 copy 一个 column row 到另外一个变量？
3、模板变量的状态。 https://grafana.com/tutorials/provision-dashboards-and-data-sources/ 。 

### Grafana 社区

1、https://community.grafana.com/t/time-delay-with-makro-unixepochgroup/49117
2、

### Grafana 授权技巧

1、https://p.sundayhey.com/w/tech/backend/arch/grafana_tutorial/
2、告警， https://p.sundayhey.com/w/tech/backend/op/grafana_alert_and_notification/ 。 
3、阿里云 SLS grafana 插件， https://github.com/mayunlei/aliyun-log-grafana-datasource-plugin/blob/master/README_CN.md 。 

4、https://github.com/aliyun/aliyun-log-grafana-datasource-plugin/blob/master/README_CN.md 。 

5、sls 相关技巧， [www.alibabacloud.com/help/zh/doc-detail/63447.htm?spm=a2c63.p38356.b99.226.38a53a00kzSVtA](https://www.alibabacloud.com/help/zh/doc-detail/63447.htm?spm=a2c63.p38356.b99.226.38a53a00kzSVtA) 。 

6、https://grafana.com/docs/grafana/v7.5/datasources/mysql/ 【需要仔细读两遍】

7、 [www.alibabacloud.com/help/zh/doc-detail/140716.htm?spm=a2c63.p38356.b99.285.16217064dAwKyj](https://www.alibabacloud.com/help/zh/doc-detail/140716.htm?spm=a2c63.p38356.b99.285.16217064dAwKyj) 。 


### 部署

1、[localhost/login](http://localhost:3000/login)
2、账号 admin , 密码 admin 
3、promethoes 用的是 PromQL 语言，
4、https://play.grafana.org/d/000000012/grafana-play-home?orgId=1 


### Readinglist 

1、https://grafana.com/tutorials/
2、https://grafana.com/tutorials/grafana-fundamentals/ fundamentals 相关。 







