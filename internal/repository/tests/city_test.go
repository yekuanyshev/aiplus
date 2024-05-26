package tests

import (
	"context"
	"os"
	"testing"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yekuanyshev/aiplus/internal/models"
	"github.com/yekuanyshev/aiplus/internal/repository"
	"github.com/yekuanyshev/aiplus/pkg/postgres"
)

type CityTestSuite struct {
	suite.Suite
	pool           *pgxpool.Pool
	cityRepository repository.City
}

func (s *CityTestSuite) TestCreate() {
	defer s.clear()

	ctx := context.Background()

	city := generateCity(s.T())
	id, err := s.cityRepository.Create(ctx, city)
	s.Require().NoError(err)
	s.Require().NotZero(id)
}

func (s *CityTestSuite) TestList() {
	defer s.clear()

	ctx := context.Background()

	var inputCities []models.City

	for i := 0; i < 3; i++ {
		inputCities = append(inputCities, generateCity(s.T()))
	}

	for i := range inputCities {
		id, err := s.cityRepository.Create(ctx, inputCities[i])
		inputCities[i].ID = id
		s.Require().NoError(err)
	}

	cities, err := s.cityRepository.List(ctx)
	s.Require().NoError(err)

	for i := 0; i < len(inputCities); i++ {
		inputCity := inputCities[i]
		city := cities[i]

		s.Require().Equal(inputCity, city)
	}
}

func (s *CityTestSuite) TestUpdate() {
	defer s.clear()

	ctx := context.Background()

	inputCity := generateCity(s.T())
	id, err := s.cityRepository.Create(ctx, inputCity)
	s.Require().NoError(err)
	inputCity.ID = id

	city := s.getCityByID(id)

	s.Require().Equal(inputCity, city)

	updatedCity := generateCity(s.T())
	updatedCity.ID = id
	err = s.cityRepository.Update(ctx, updatedCity)
	s.Require().NoError(err)

	city = s.getCityByID(id)
	s.Require().Equal(updatedCity, city)
}

func (s *CityTestSuite) getCityByID(id int64) models.City {
	var (
		ctx   = context.Background()
		query = `
			SELECT id, name
			FROM city
			WHERE id = $1
			LIMIT 1;
		`
		city models.City
	)

	err := pgxscan.Get(ctx, s.pool, &city, query, id)
	s.Require().NoError(err)
	return city
}

func (s *CityTestSuite) clear() {
	ctx := context.Background()
	query := `DELETE FROM city WHERE id > 0`
	_, err := s.pool.Exec(ctx, query)
	s.Require().NoError(err)
}

func TestCityTestSuite(t *testing.T) {
	pgDSN := os.Getenv("PG_DSN")
	require.NotEmpty(t, pgDSN, "empty PG_DSN")

	pool, err := postgres.Connect(context.Background(), pgDSN)
	require.NoError(t, err)

	suite.Run(t, &CityTestSuite{
		pool:           pool,
		cityRepository: repository.NewCity(pool),
	})
}
