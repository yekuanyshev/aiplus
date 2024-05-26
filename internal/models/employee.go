package models

type Employee struct {
	ID         int64   `db:"id"`
	Phone      string  `db:"phone"`
	FirstName  string  `db:"first_name"`
	LastName   string  `db:"last_name"`
	MiddleName *string `db:"middle_name"`
	CityID     int64   `db:"city_id"`
}
