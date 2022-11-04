package main

import (
	"os"
	"os/signal"

	"github.com/smirzoavliyoev/xmtask/pkg/logger"
	"github.com/smirzoavliyoev/xmtask/pkg/nats/config"
	"github.com/smirzoavliyoev/xmtask/pkg/nats/connection"
)

func main() {

	// collect all modules here ->
	// define all dependencies and inject to each other
	// dependencies should be interfaces
	logger := logger.NewLogger()
	conn := connection.NewConn(config.Config{
		ClusterID: "some",
		ClientID:  "some",
	})
	defer conn.Close()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	//serve server
	//TODO:: gracefully shutdown server
	<-ch
	logger.Sync()

}
