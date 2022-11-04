package company

import (
	"errors"

	"github.com/smirzoavliyoev/xmtask/pkg/repositories/companies"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type Company struct {
	// TODO:: describe dependencies using interfaces
	repo   companies.CompanyRepo
	logger *zap.SugaredLogger
}

func NewCompanyService(repo *companies.CompanyRepo, logger *zap.SugaredLogger) *Company {
	return &Company{
		repo:   *repo,
		logger: logger,
	}
}

func (c *Company) GetCompany(f companies.CompanyFilter) ([]companies.Company, error) {
	companies, err := c.repo.GetCompany(f)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.logger.Error("error while trying to fetch data from repository")
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		c.logger.Error("record not found", err)
		return nil, ErrNotFound
	}

	return companies, nil
}

func (c *Company) Create(CompanyDto companies.Company) error {
	err := c.repo.Create(CompanyDto)
	if err != nil {
		c.logger.Error("error while trying to create data from repository")

		return err
	}

	return nil
}

func (c *Company) Update(CompanyFilter companies.CompanyFilter) error {
	err := c.repo.Update(CompanyFilter)
	if err != nil {
		c.logger.Error("error while trying to update data from repository")

		return err
	}

	return nil
}

func (c *Company) Delete(ids ...int) error {
	err := c.repo.Delete(ids...)
	if err != nil {
		c.logger.Error("error while trying to delete data from repository")

		return err
	}
	return nil
}
