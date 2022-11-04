package handlers

import (
	company "github.com/smirzoavliyoev/xmtask/internal/companyservice"
	"go.uber.org/zap"
)

type Handlers struct {
	// use interface instead of directly dependencies
	// describe interfaces in client side
	companyService *company.Company
	logger         *zap.SugaredLogger
}

func NewHandlers(companyService *company.Company, logger *zap.SugaredLogger) *Handlers {
	return &Handlers{
		companyService: companyService,
		logger:         logger,
	}
}
