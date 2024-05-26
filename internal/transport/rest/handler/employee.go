package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yekuanyshev/aiplus/internal/service"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler/form"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler/response"
)

type employee struct {
	service service.Manager
}

func NewEmployee(service service.Manager) *employee {
	return &employee{
		service: service,
	}
}

func (h *employee) Create(c *fiber.Ctx) error {
	var employeeCreate form.EmployeeCreate
	err := c.BodyParser(&employeeCreate)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	id, err := h.service.Employee().Create(c.UserContext(), employeeCreate.ToModel())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, response.EmployeeCreate{
		ID: id,
	})
}

func (h *employee) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	employee, err := h.service.Employee().GetByID(c.UserContext(), int64(id))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, response.Employee{
		Phone:      employee.Phone,
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		MiddleName: employee.MiddleName,
		CityID:     employee.CityID,
	})
}
