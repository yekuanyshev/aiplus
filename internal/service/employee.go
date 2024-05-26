package service

import (
	"context"

	"github.com/yekuanyshev/aiplus/internal/models"
	"github.com/yekuanyshev/aiplus/internal/repository"
)

type employee struct {
	repository repository.Manager
}

func NewEmployee(repository repository.Manager) *employee {
	return &employee{
		repository: repository,
	}
}

func (srv *employee) Create(ctx context.Context, employee models.Employee) (id int64, err error) {
	return srv.repository.Employee().Create(ctx, employee)
}

func (srv *employee) GetByID(ctx context.Context, id int64) (employee models.Employee, err error) {
	return srv.repository.Employee().ByID(ctx, id)
}
