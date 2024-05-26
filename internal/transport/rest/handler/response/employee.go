package response

type EmployeeCreate struct {
	ID int64 `json:"id"`
}

type Employee struct {
	Phone      string  `json:"phone"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MiddleName *string `json:"middleName"`
	CityID     int64   `json:"cityID"`
}
