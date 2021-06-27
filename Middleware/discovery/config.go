package discovery

import (
	"go-common/library/log"
	// "go-common/library/net/http"
	http "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// Env is discovery env.
type Env struct {
	Region    string
	Zone      string
	Host      string
	DeployEnv string
}

// Config config.
type Config struct {
	Nodes         []string
	Zones         map[string][]string
	HTTPServer    *http.ServerConfig
	// HTTPClient    *http.ClientConfig
	Env           *Env
	Log           *log.Config

	// 这里为什么是 byte 的形式
	Scheduler     []byte
	EnableProtect bool
}

