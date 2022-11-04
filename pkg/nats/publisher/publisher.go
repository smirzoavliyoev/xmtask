package publisher

import "github.com/nats-io/stan.go"

type Publisher struct {
	sc stan.Conn
}

func NewPublisher(sc stan.Conn) *Publisher {
	return &Publisher{
		sc: sc,
	}
}

func (p *Publisher) Publish(clusterName string, message []byte) error {
	return p.sc.Publish(clusterName, message)
}
