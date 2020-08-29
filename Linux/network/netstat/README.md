### netstat 的常见使用方法
1. https://www.cnblogs.com/ggjucheng/archive/2012/01/08/2316661.HTML 
2. 

### 想在通常用 SS 命令替代 netstat 
1. https://linux.die.net/man/8/netstat

### 查看是否有连接溢出
1. netstat -s | grep LISTEN

### 参考连接
1. https://qcrao.com/2019/01/16/dive-into-three-way-handshake/
2. HTTP 分析， https://qcrao.com/2019/01/18/dive-into-http/
3. backlog, https://developpaper.com/understand-the-handling-of-tcp-handshake-by-the-server-and-what-is-backlog/
4. 利用 netstat 监控大量 ESTABLISHED 连接书和 TIME_WAIT 数量， https://blog.nowcoder.net/n/758c3a57e80940eba894d5e5172a4186