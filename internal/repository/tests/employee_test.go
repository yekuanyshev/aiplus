package tests

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yekuanyshev/aiplus/internal/repository"
	"github.com/yekuanyshev/aiplus/pkg/postgres"
)

type EmployeeTestSuite struct {
	suite.Suite
	pool            *pgxpool.Pool
	cityManager     repository.City
	employeeManager repository.Employee
}

func (s *EmployeeTestSuite) TestCreate() {
	defer s.clear()

	ctx := context.Background()

	cityID := s.createCity()
	employee := generateEmployee(cityID)
	employeeID, err := s.employeeManager.Create(ctx, employee)
	s.Require().NoError(err)
	s.Require().NotEmpty(employeeID)
}

func (s *EmployeeTestSuite) TestByID() {
	defer s.clear()

	ctx := context.Background()
	cityID := s.createCity()
	createEmployee := generateEmployee(cityID)

	id, err := s.employeeManager.Create(ctx, createEmployee)
	s.Require().NoError(err)
	createEmployee.ID = id

	employee, err := s.employeeManager.ByID(ctx, id)
	s.Require().NoError(err)

	s.Require().Equal(createEmployee, employee)
}

func (s *EmployeeTestSuite) createCity() int64 {
	ctx := context.Background()

	city := generateCity(s.T())
	cityID, err := s.cityManager.Create(ctx, city)
	s.Require().NoError(err)
	s.Require().NotEmpty(cityID)
	return cityID
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

	pool, err := postgres.Connect(context.Background(), pgDSN)
	require.NoError(t, err)

	suite.Run(t, &EmployeeTestSuite{
		pool:            pool,
		cityManager:     repository.NewCity(pool),
		employeeManager: repository.NewEmployee(pool),
	})
}
