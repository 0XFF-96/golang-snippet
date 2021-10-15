# sqlx 文档

### these new semantics:
1.  ad-hoc query execution
2. The ? query placeholders, called bindvars internally, are important
3. 

### illustrated guide to sqlx
1. http://jmoiron.github.io/sqlx/ (有几个骚技巧可以学一下)
2. https://pkg.go.dev/database/sql

### 一些关于事务的建议

给http请求设置超时时间，此为最后的保命手段，其它所有服务都应有合理设置
事务的初始化也要设置超时时间，所有使用事务的方法都应该确保设置有合理的超时时间
事务执行应该独立一个连接池，不与其它查询共用
要好好看文档，事务开启后要使用*Tx完成所有操作，文档中有明确的说明，当前的用法是无法实现事务控制预期的

