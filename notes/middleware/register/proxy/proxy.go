package proxy


import (
	// grpclb "code.aliyun.com/xchat/xgo/lb"
	// "code.aliyun.com/xchat/xgo/registry/consul"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// Proxy service proxy
type Proxy struct {
	Conn *grpc.ClientConn
}

// NewProxy new service proxy
func NewProxy(serviceName string, address string) (*Proxy, error) {
	// resolver := consul.NewResolver(serviceName, address)
	// balancer := grpclb.NewBalancer(resolver, grpclb.NewRoundRobinSelector())

	unaryInterceptors := grpc_middleware.ChainUnaryClient(
		grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())),
	)

	conn, err := grpc.Dial(
		"",
		grpc.WithInsecure(),
		// grpc.WithBalancer(balancer),
		grpc.WithUnaryInterceptor(unaryInterceptors),
	)
	if err != nil {
		return nil, err
	}
	return &Proxy{
		Conn: conn,
	}, nil
}

// Close close connection
func (p *Proxy) Close() {
	p.Conn.Close()
}