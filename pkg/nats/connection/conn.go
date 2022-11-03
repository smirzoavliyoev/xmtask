package connection

import (
	"github.com/nats-io/stan.go"
	"github.com/smirzoavliyoev/xmtask/pkg/nats/config"
)

func NewConn(config config.Config) stan.Conn {
	var err error
	conn, err := stan.Connect(config.ClusterID, config.ClientID)
	if err != nil {
		panic(err)
	}
	return conn
}
