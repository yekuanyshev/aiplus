package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler/form"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler/response"
	"github.com/yekuanyshev/aiplus/pkg/postgres"
)

type EmployeeTestSuite struct {
	suite.Suite

	pool     *pgxpool.Pool
	basePath string
}

func (s *EmployeeTestSuite) TestCreate() {
	defer s.clear()

	cityID := s.createCity("Almaty")
	id := s.createEmployee(cityID)
	s.Require().NotZero(id)
}

func (s *EmployeeTestSuite) TestGetByID() {
	defer s.clear()

	var (
		cityID          = s.createCity("Almaty")
		id              = s.createEmployee(cityID)
		url             = fmt.Sprintf(s.basePath+"/employee/%d", id)
		createdEmployee struct {
			Success bool              `json:"success"`
			Data    response.Employee `json:"data"`
		}
	)

	response, err := http.Get(url)
	s.Require().NoError(err)
	s.Require().Equal(response.StatusCode, fiber.StatusOK)

	err = json.NewDecoder(response.Body).Decode(&createdEmployee)
	s.Require().NoError(err)
	s.Require().True(createdEmployee.Success)
	s.Require().NotEmpty(createdEmployee.Data)
}

func (s *EmployeeTestSuite) createEmployee(cityID int64) int64 {
	var (
		url            = s.basePath + "/employee"
		employeeCreate = form.EmployeeCreate{
			Phone:     faker.Phonenumber(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			CityID:    cityID,
		}
		buff            = bytes.NewBuffer(nil)
		createdEmployee struct {
			Success bool                    `json:"success"`
			Data    response.EmployeeCreate `json:"data"`
		}
	)

	err := json.NewEncoder(buff).Encode(employeeCreate)
	s.Require().NoError(err)

	response, err := http.Post(url, "application/json", buff)
	s.Require().NoError(err)
	s.Require().Equal(response.StatusCode, fiber.StatusOK)

	err = json.NewDecoder(response.Body).Decode(&createdEmployee)
	s.Require().NoError(err)
	s.Require().True(createdEmployee.Success)

	return createdEmployee.Data.ID
}

func (s *EmployeeTestSuite) createCity(name string) int64 {
	var (
		url        = s.basePath + "/city"
		cityCreate = form.CityCreate{
			Name: name,
		}
		buff        = bytes.NewBuffer(nil)
		createdCity struct {
			Success bool          `json:"success"`
			Data    response.City `json:"data"`
		}
	)

	err := json.NewEncoder(buff).Encode(cityCreate)
	s.Require().NoError(err)

	response, err := http.Post(url, "application/json", buff)
	s.Require().NoError(err)
	s.Require().Equal(response.StatusCode, fiber.StatusOK)

	err = json.NewDecoder(response.Body).Decode(&createdCity)
	s.Require().NoError(err)
	s.Require().True(createdCity.Success)
	return createdCity.Data.ID
}

func (s *EmployeeTestSuite) clear() {
	ctx := context.Background()
	query := `
		DELETE FROM employee WHERE id > 0;
		DELETE FROM city WHERE id > 0;
	`
	_, err := s.pool.Exec(ctx, query)
	s.Require().NoError(err)
}

func TestEmployeeTestSuite(t *testing.T) {
	pgDSN := os.Getenv("PG_DSN")
	require.NotEmpty(t, pgDSN, "empty PG_DSN")

	basePath := os.Getenv("BASE_PATH")
	require.NotEmpty(t, pgDSN, "empty BASE_PATH")

	pool, err := postgres.Connect(context.Background(), pgDSN)
	require.NoError(t, err)

	suite.Run(t, &CityTestSuite{
		pool:     pool,
		basePath: basePath,
	})
}
