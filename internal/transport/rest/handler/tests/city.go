package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler/form"
	"github.com/yekuanyshev/aiplus/pkg/postgres"
)

type CityTestSuite struct {
	suite.Suite

	pool     *pgxpool.Pool
	basePath string
}

func (s *CityTestSuite) TestCreate() {
	defer s.clear()

	id := s.createCity("Almaty")
	s.Require().NotZero(id)
}

func (s *CityTestSuite) TestList() {
	defer s.clear()

	var (
		url         = s.basePath + "/city"
		inputCities = []form.CityCreate{
			{Name: "Almaty"},
			{Name: "Astana"},
			{Name: "Karaganda"},
		}
		cities struct {
			Success bool `json:"success"`
			Data    []struct {
				ID   int64  `json:"id"`
				Name string `json:"name"`
			} `json:"data"`
		}
	)

	for i := 0; i < len(inputCities); i++ {
		s.createCity(inputCities[i].Name)
	}

	response, err := http.Get(url)
	s.Require().NoError(err)
	s.Require().Equal(response.StatusCode, fiber.StatusOK)

	err = json.NewDecoder(response.Body).Decode(&cities)
	s.Require().NoError(err)
	s.Require().True(cities.Success)
	s.Require().NotEmpty(cities.Data)

	for i := range cities.Data {
		inputCity := inputCities[i]
		city := cities.Data[i]
		s.Require().Equal(inputCity.Name, city.Name)
	}
}

func (s *CityTestSuite) TestUpdate() {
	defer s.clear()

	var (
		id          = s.createCity("Almaty")
		url         = fmt.Sprintf(s.basePath+"/city/%d", id)
		updatedCity = form.CityUpdate{
			Name: "Astana",
		}
		buff = bytes.NewBuffer(nil)
	)

	err := json.NewEncoder(buff).Encode(updatedCity)
	s.Require().NoError(err)

	request, err := http.NewRequest(http.MethodPut, url, buff)
	s.Require().NoError(err)

	response, err := http.DefaultClient.Do(request)
	s.Require().NoError(err)
	s.Require().Equal(response.StatusCode, fiber.StatusOK)
}

func (s *CityTestSuite) createCity(name string) int64 {
	var (
		url        = s.basePath + "/city"
		cityCreate = form.CityCreate{
			Name: name,
		}
		buff        = bytes.NewBuffer(nil)
		createdCity struct {
			Success bool `json:"success"`
			Data    struct {
				ID int64 `json:"id"`
			} `json:"data"`
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

func (s *CityTestSuite) clear() {
	ctx := context.Background()
	query := `DELETE FROM city WHERE id > 0`
	_, err := s.pool.Exec(ctx, query)
	s.Require().NoError(err)
}

func TestCityTestSuite(t *testing.T) {
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
