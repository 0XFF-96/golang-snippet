### Wait-For-ready 

// WaitForReady configures the action to take when an RPC is attempted on broken
// connections or unreachable servers. If waitForReady is false, the RPC will fail
// immediately. Otherwise, the RPC client will block the call until a
// connection is available (or the call is canceled or times out) and will
// retry the call if it fails due to a transient error.  gRPC will not retry if
// data was written to the wire unless the server indicates it did not process
// the data.  Please refer to
// https://github.com/grpc/grpc/blob/master/doc/wait-for-ready.md.

### Retry 

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

- 这个例子没有成功，期望是调用一次 client 代码，然后能够检测到发送 3 次包的。 
- 

### Reflection 

- https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md#grpc-cli.
- There are multiple existing reflection clients.
- 利用反射的例子： https://github.com/fullstorydev/grpcurl 