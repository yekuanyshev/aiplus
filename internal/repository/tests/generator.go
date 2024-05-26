package tests

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
	"github.com/yekuanyshev/aiplus/internal/models"
)

func generateCity(t *testing.T) models.City {
	city := models.City{}
	err := faker.FakeData(&city)
	require.NoError(t, err)
	return city
}

func generateEmployee(cityID int64) models.Employee {
	middleName := faker.FirstName()
	employee := models.Employee{
		Phone:      faker.Phonenumber(),
		FirstName:  faker.FirstName(),
		LastName:   faker.LastName(),
		MiddleName: &middleName,
		CityID:     cityID,
	}

	return employee
}
