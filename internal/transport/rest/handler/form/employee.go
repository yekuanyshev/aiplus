package form

import "github.com/yekuanyshev/aiplus/internal/models"

type EmployeeCreate struct {
	Phone      string  `json:"phone"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MiddleName *string `json:"middleName"`
	CityID     int64   `json:"cityID"`
}

func (e EmployeeCreate) ToModel() models.Employee {
	return models.Employee{
		Phone:      e.Phone,
		FirstName:  e.FirstName,
		LastName:   e.LastName,
		MiddleName: e.MiddleName,
		CityID:     e.CityID,
	}
}
