package go_proxy

import (
	"net"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// ProxyHappySuite tests the "happy" path of handling:
type ProxyHappySuite struct {
	suite.Suite// that everything works in absence of connection issues.


	serverListener   net.Listener
	server           *grpc.Server
	proxyListener    net.Listener
	proxy            *grpc.Server
	serverClientConn *grpc.ClientConn

	client     *grpc.ClientConn
	// testClient pb.TestServiceClient
}

// 在这里两个函数中
// 把一些其他测试用例也需要用到的初始化数据，放在这里
// SetUpSuite 在所有测试用例运行前，需要运行的函数
func (s *ProxyHappySuite) SetupSuite() {
}

// 在所有测试用例运行后，需要运行的函数
func (s *ProxyHappySuite) TearDownSuite() {
}
