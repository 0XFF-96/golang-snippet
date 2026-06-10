### 什么是 NameResolver 
A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.

- NameResolver 是在 client 端 Dial 的时候作用的
-  