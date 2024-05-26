package service

import (
	"context"

	"github.com/yekuanyshev/aiplus/internal/models"
	"github.com/yekuanyshev/aiplus/internal/repository"
)

var (
	_ Manager  = &manager{}
	_ City     = &city{}
	_ Employee = &employee{}
)

type (
	Manager interface {
		City() City
		Employee() Employee
	}

	City interface {
		Create(ctx context.Context, city models.City) (id int64, err error)
		List(ctx context.Context) (cities []models.City, err error)
		Update(ctx context.Context, city models.City) (err error)
	}

	Employee interface {
		Create(ctx context.Context, employee models.Employee) (id int64, err error)
		GetByID(ctx context.Context, id int64) (employee models.Employee, err error)
	}

	manager struct {
		city     *city
		employee *employee
	}
)

func NewManager(repository repository.Manager) Manager {
	return &manager{
		city:     NewCity(repository),
		employee: NewEmployee(repository),
	}
}

func (m *manager) City() City {
	return m.city
}

func (m *manager) Employee() Employee {
	return m.employee
}
