package companies

type Company struct {
	Name    string
	Code    string
	Country string
	Website string
	Phone   string
}

type CompanyFilter struct {
	ID      uint
	Name    string
	Code    string
	Country string
	Website string
	Phone   string
}
