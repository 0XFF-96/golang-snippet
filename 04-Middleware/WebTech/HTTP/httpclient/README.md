# http client 

## 前沿
在日常开发中，我们往往需要对接其他的第三方服务。例如，支付宝支付、微信支付、厂商推送等等。
在这些需求中有一个共同的特点，就是你需要利用 http 请求，调用第三方接口，完成相应的操作。

通常我们会直接利用 golang 标准库提供的函数，完成 PUT、GET、DELETE、POST 等请求。
有一个更好的做法是：抽像出一个 http client 中间层给我们业务逻辑调用。

这样做的好处是：
1、易扩展。
2、可定制化。
3、业务逻辑分离
4、提供不同的能力：可以定制超时时间、实现是否超时功能、实现鉴权功能
5、可以复用
6、以后有调用外部 http 服务的代码，都收敛到一起。

## 代码解析
（略, code extracted from hua-wei SDK）

##参考

* [支付宝 golang SDK](https://github.com/smartwalle/alipay)
* [支付宝 golang SDK2](https://github.com/ascoders/alipay/blob/master/alipay.go)
* [微信支付 golang SDK](https://github.com/smartwalle/wxpay)
* [微信支付 golang SDK](https://github.com/objcoding/wxpay)
* [华为 push SDK](...)

## what's next:
1、增加相应测试
2、更好地打日志 log 
3、看看还有没有其他可以改进的点？
