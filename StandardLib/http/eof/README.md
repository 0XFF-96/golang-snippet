

### req.Close 的相关东西

	// Close indicates whether to close the connection after
	// replying to this request (for servers) or after sending this
	// request and reading its response (for clients).
	//
	// For server requests, the HTTP server handles this automatically
	// and this field is not needed by Handlers.
	//
	// For client requests, setting this field prevents re-use of
	// TCP connections between requests to the same hosts, as if
	// Transport.DisableKeepAlives were set.
	
[For server requests]: 什么是 server requests ? handle request 这个参数有没有用？

[For client requests]: 什么是 client requests ?

req.Close 用到的地方:
```
		if resp.Close || rc.req.Close || resp.StatusCode <= 199 || bodyWritable {
			// Don't do keep-alive on error if either party requested a close
			// or we get an unexpected informational (1xx) response.
			// StatusCode 100 is already handled above.
			alive = false
		}
```

这两个的意思是指 incommingCtx, outgoingCtx 的区别 
	
	

### 为什么只有 upush 有这个问题，其他没有？
- 跟 upush 服务器的行为有关系，如下图。
```
The problem that I ran into is that the server is responding with Connection: 
Keep-Alive in the response header and then immediately closing the connection
```


And the req.Close = true, solution works 
because it prevents the connection from being re-used


The second GET request fails with "can't write HTTP request on broken connection" or
"EOF".