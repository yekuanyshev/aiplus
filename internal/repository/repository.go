package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yekuanyshev/aiplus/internal/models"
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
		ByID(ctx context.Context, id int64) (employee models.Employee, err error)
	}

	manager struct {
		city     *city
		employee *employee
	}
)

func NewManager(pool *pgxpool.Pool) Manager {
	return &manager{
		city:     NewCity(pool),
		employee: NewEmployee(pool),
	}
}

func (m *manager) City() City {
	return m.city
}

func (m *manager) Employee() Employee {
	return m.employee
}
