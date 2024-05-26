package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yekuanyshev/aiplus/internal/models"
	"github.com/yekuanyshev/aiplus/internal/service"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler/form"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler/response"
)

type city struct {
	service service.Manager
}

func NewCity(service service.Manager) *city {
	return &city{
		service: service,
	}
}

func (h *city) Create(c *fiber.Ctx) error {
	var cityCreate form.CityCreate
	err := c.BodyParser(&cityCreate)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	city := models.City{
		Name: cityCreate.Name,
	}

	id, err := h.service.City().Create(c.UserContext(), city)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, response.CityCreate{
		ID: id,
	})
}

func (h *city) List(c *fiber.Ctx) error {
	cities, err := h.service.City().List(c.UserContext())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	var responseCity []response.City
	for _, city := range cities {
		responseCity = append(responseCity, response.City{
			ID:   city.ID,
			Name: city.Name,
		})
	}

	return response.Success(c, responseCity)
}

func (h *city) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	var cityUpdate form.CityUpdate
	err = c.BodyParser(&cityUpdate)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	city := models.City{
		ID:   int64(id),
		Name: cityUpdate.Name,
	}

	err = h.service.City().Update(c.UserContext(), city)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, nil)
}
