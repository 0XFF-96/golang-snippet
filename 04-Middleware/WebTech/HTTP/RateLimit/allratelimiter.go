package RateLimit

import (
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/google/netstack/tcpip/stack"
	"github.com/jaegertracing/jaeger-client-go"
	"github.com/jaegertracing/jaeger-client-go/utils"
	"k8s.io/kubernetes/pkg/controller/podautoscaler"
	"github.com/sundayfun/daycam-server/pkg/ratelimit"
)


func testNew(){
	ratelimit.NewLimiter(nil)
	blademaster.NewRateLimiter(nil)
	jaeger.NewRateLimitingSampler(0)
	utils.NewRateLimiter(0, 0)
	podautoscaler.NewDefaultHPARateLimiter(0)
	stack.NewICMPRateLimiter()
	podautoscaler2.NewDefaultHPARateLimiter(0)
}

