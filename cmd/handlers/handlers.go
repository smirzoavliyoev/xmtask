package handlers

import company "github.com/smirzoavliyoev/xmtask/internal/companyservice"

type Handlers struct {
	companyService *company.Company
}

func NewHandlers(companyService *company.Company) *Handlers {
	return &Handlers{
		companyService: companyService,
	}
}
