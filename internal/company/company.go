package company

import "github.com/smirzoavliyoev/xmtask/pkg/repositories/companies"

type Company struct {
	repo companies.CompanyRepo
}

func NewCompanyService(repo *companies.CompanyRepo) *Company {
	return &Company{
		repo: *repo,
	}
}

func (c *Company) GetCompany(f companies.CompanyFilter) {

}
