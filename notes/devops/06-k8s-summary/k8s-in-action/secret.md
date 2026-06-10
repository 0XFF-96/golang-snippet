## 什么是 secret ?

敏感的配置信息，例如密钥、证书等则通常由后者来配置。

在本实验场景中，我们会带你学习使用Secret资源对象为Pod注入敏感配置信息，这包括

1、Secret 资源
2、基于存储卷接口使用 Secret 

分类
1、generic：基于本地文件、目录或字面量值创建的Secret，一般用来存储密码、密钥、信息、证书等数据
2、docker-registry：专用于认证到Docker Registry的Secret，以使用私有容器镜像
3、tls：基于指定的公钥/私钥对创建TLS Secret，专用于TLS通信中；指定公钥和私钥必须事先存在，公钥证书必须采用PEM编码，且应该与指定的私钥相匹配。