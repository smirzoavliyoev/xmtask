package companies

import (
	"time"

	"gorm.io/gorm"
)

type CompanyRepo struct {
	conn *gorm.DB
}

func NewCompanyRepo(conn *gorm.DB) *CompanyRepo {
	return &CompanyRepo{
		conn: conn,
	}
}

type company struct {
	ID      uint
	Name    string
	Code    string
	Country string
	Website string
	Phone   string

	createdAt time.Time
	updatedAt time.Time
}

type DatabaseConnection interface {
}

func (c *CompanyRepo) GetCompany(f CompanyFilter) ([]Company, error) {
	result := make([]Company, 0)
	rows, err := c.conn.Where(&f).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i Company
		err := rows.Scan(&i.Name, &i.Code, &i.Country, &i.Website, &i.Phone)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

func (c *CompanyRepo) Create(companyDto Company) error {

	var companyDB = company{
		Name:    companyDto.Name,
		Code:    companyDto.Code,
		Country: companyDto.Country,
		Website: companyDto.Website,
		Phone:   companyDto.Phone,

		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	res := c.conn.Create(&companyDB)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyRepo) Update(companyFilter CompanyFilter) error {
	res := c.conn.Save(&companyFilter)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (c *CompanyRepo) Delete(ids ...int) error {
	res := c.conn.Delete(&company{}, ids)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
