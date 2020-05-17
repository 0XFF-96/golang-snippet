package redis

import (
	"github.com/CodisLabs/codis/pkg/utils/redis"
	redigo "github.com/garyburd/redigo/redis"
)

func main() {
	p := redis.Pool{}
	p.Close()
	client, err := redigo.Dial("", "")
	if err != nil {
	}
	client.Close()
}
