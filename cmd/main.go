package main

import (
	"os"
	"os/signal"

	"github.com/smirzoavliyoev/xmtask/pkg/nats/config"
	"github.com/smirzoavliyoev/xmtask/pkg/nats/connection"
)

func main() {

	// collect all modules here ->
	// define all dependencies and inject to each other

	conn := connection.NewConn(config.Config{
		ClusterID: "some",
		ClientID:  "some",
	})
	defer conn.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)

	<-ch
}
