package service

import (
	"context"

	"github.com/yekuanyshev/aiplus/internal/models"
	"github.com/yekuanyshev/aiplus/internal/repository"
)

type city struct {
	repository repository.Manager
}

func NewCity(repository repository.Manager) *city {
	return &city{
		repository: repository,
	}
}

func (srv *city) Create(ctx context.Context, city models.City) (id int64, err error) {
	return srv.repository.City().Create(ctx, city)
}

func (srv *city) List(ctx context.Context) (cities []models.City, err error) {
	return srv.repository.City().List(ctx)
}

func (srv *city) Update(ctx context.Context, city models.City) (err error) {
	return srv.repository.City().Update(ctx, city)
}
