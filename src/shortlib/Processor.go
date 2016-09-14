package shortlib

import(
	"net/http"
)

type Processor interface {
	ProcessRequest(method, request_url string, params map[string]string, body []byte, w http.ResponseWriter, r *http.Request) error
}

type BaseProcessor struct {
	RedisCli *RedisAdaptor
	Configure *Configure
	HostName string
	Lru		 *LRU
	CountFunction CreateCountFunc
}