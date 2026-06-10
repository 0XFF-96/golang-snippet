package proxy

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"

	//"github.com/google/uuid"
	//"github.com/syncthing/syncthing/lib/rc"

	// "code.aliyun.com/xchat/xgo/registry/consul"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	// uuid "github.com/satori/go.uuid"
	jaegerc "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
)
// ServerConfig server config

type ServerConfig struct {
	Address   string `yaml:"address"`
	Name      string `yaml:"name"`
	Timeout   int    `yaml:"timeout"`
	LogLevel  string `yaml:"logLevel"`
	DockerEnv string `yaml:"dockerEnv"`
}

// GrpcServer include grpc server, registry
type GrpcServer struct {
	Server   *grpc.Server
	// registry *consul.Registry
	sc       ServerConfig
	closer   io.Closer
}

// GetIPAddress get ip address
func GetIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// NewGrpcServer create new grpc server
func NewGrpcServer(
	sc *ServerConfig,
	// rc *consul.RegistryConfig,
	jc *jaegerc.Configuration,
) (*GrpcServer, error) {
	ap := strings.Split(sc.Address, ":")
	if len(ap) < 2 {
		return nil, errors.New("server address is error")
	}
	address := ap[0]
	if address == "" {
		address = "0.0.0.0"
	}
	if sc.Name == "" {
		return nil, errors.New("server name is empty")
	}
	port, err := strconv.Atoi(ap[1])
	if err != nil {
		return nil, fmt.Errorf("invalid server port %s", ap[1])
	}
	if port == 0 {
		return nil, errors.New("invalid server port 0")
	}

	//rc.Name = sc.Name
	//if address == "0.0.0.0" {
	//	rc.NodeAddress = GetIPAddress()
	//} else {
	//	rc.NodeAddress = address
	//}
	//rc.NodePort = port
	//if rc.TTL <= 0 {
	//	rc.TTL = 5
	//}
	//uuidv1 := uuid.NewV1()
	//rc.ID = uuidv1.String()
	//registry, err := consul.NewRegistry(rc)
	//if err != nil {
	//	return nil, err
	//}
	//
	//_, closer, err := NewJaegerTracer(sc.Name, jc)
	//if err != nil {
	//	fmt.Printf("NewJaegerTracer() error %v", err)
	//}
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())),
		),
	)

	return &GrpcServer{
		Server:   server,
		// registry: registry,
		sc:       *sc,
		// closer:   closer,
	}, nil
}

// Run start grpc server
func (s *GrpcServer) Run() {
	fmt.Println("app start")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Printf("starting service: %s - %s\n", s.sc.Name, s.sc.Address)
		listener, err := net.Listen("tcp", s.sc.Address)
		if err != nil {
			panic(fmt.Errorf("%v", err))
		}
		s.Server.Serve(listener)
		wg.Done()
		fmt.Println("service stopped")
	}()

	wg.Add(1)
	go func() {
		fmt.Println("registering service")
		// 进行服务注册的地方
		// err := s.registry.Register()
/*		if err != nil {
			fmt.Println(err)
		}*/
		wg.Done()
		fmt.Println("service deregistered")
	}()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-sigs
		fmt.Println("receive signal: " + sig.String())
		done <- true
	}()
	<-done
	fmt.Println("stopping service")
	s.Server.Stop()
	fmt.Println("deregistering service")

	// 服务停止后，解除注册的地方
	// s.registry.Deregister()

	wg.Wait()
	fmt.Print("app stopped")
}

