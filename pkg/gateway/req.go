package gateway

import (
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

var once = sync.Once{}
var client *req.Client

func NewClient() *req.Client {
	once.Do(newClient)
	return client
}

func newClient() {
	client = req.C().SetTimeout(5 * time.Second)
}
