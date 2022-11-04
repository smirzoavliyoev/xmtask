package main

import (
	"os"
	"os/signal"

	"github.com/smirzoavliyoev/xmtask/cmd/handlers"
	company "github.com/smirzoavliyoev/xmtask/internal/companyservice"
	"github.com/smirzoavliyoev/xmtask/pkg/logger"
	"github.com/smirzoavliyoev/xmtask/pkg/nats/config"
	"github.com/smirzoavliyoev/xmtask/pkg/nats/connection"
	"github.com/smirzoavliyoev/xmtask/pkg/nats/publisher"
	"github.com/smirzoavliyoev/xmtask/pkg/repositories"
	"github.com/smirzoavliyoev/xmtask/pkg/repositories/companies"
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
	db := repositories.NewDB()
	companiesRepository := companies.NewCompanyRepo(db)
	companyService := company.NewCompanyService(companiesRepository, logger)

	pub := publisher.NewPublisher(conn)

	handlersService := handlers.NewHandlers(companyService, logger, pub)

	go func() {
		handlers.NewRouter(handlersService)
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	//serve server
	//TODO:: gracefully shutdown server
	<-ch
	conn.Close()
	logger.Sync()
}
