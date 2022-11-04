package handlers

import (
	company "github.com/smirzoavliyoev/xmtask/internal/companyservice"
	"github.com/smirzoavliyoev/xmtask/pkg/nats/publisher"
	"go.uber.org/zap"
)

type Handlers struct {
	// use interface instead of directly dependencies
	// describe interfaces in client side
	companyService *company.Company
	logger         *zap.SugaredLogger
	natsPub        *publisher.Publisher
}

func NewHandlers(companyService *company.Company, logger *zap.SugaredLogger, pub *publisher.Publisher) *Handlers {
	return &Handlers{
		companyService: companyService,
		logger:         logger,
		natsPub:        pub,
	}
}
