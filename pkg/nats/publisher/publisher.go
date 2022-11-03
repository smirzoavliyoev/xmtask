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

func (p *Publisher) Publish(clusterName string, message string) error {
	return p.sc.Publish(clusterName, []byte(message))
}
