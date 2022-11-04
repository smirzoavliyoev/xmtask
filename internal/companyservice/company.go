package company

import (
	"errors"

	"github.com/smirzoavliyoev/xmtask/pkg/repositories/companies"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type Company struct {
	repo companies.CompanyRepo
}

func NewCompanyService(repo *companies.CompanyRepo) *Company {
	return &Company{
		repo: *repo,
	}
}

func (c *Company) GetCompany(f companies.CompanyFilter) ([]companies.Company, error) {
	companies, err := c.repo.GetCompany(f)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}

	return companies, nil
}

func (c *Company) Create(CompanyDto companies.Company) error {
	err := c.repo.Create(CompanyDto)
	if err != nil {
		return err
	}

	return nil
}

func (c *Company) Update(CompanyFilter companies.CompanyFilter) error {
	err := c.repo.Update(CompanyFilter)
	if err != nil {
		return err
	}

	return nil
}

func (c *Company) Delete(ids ...int) error {
	err := c.repo.Delete(ids...)
	if err != nil {
		return err
	}
	return nil
}
