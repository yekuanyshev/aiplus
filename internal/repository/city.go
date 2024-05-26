package repository

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yekuanyshev/aiplus/internal/models"
)

type city struct {
	pool         *pgxpool.Pool
	queryBuilder squirrel.StatementBuilderType
}

func NewCity(pool *pgxpool.Pool) *city {
	return &city{
		pool:         pool,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (repo *city) Create(ctx context.Context, city models.City) (id int64, err error) {
	query, args, err := repo.queryBuilder.
		Insert("city").
		Columns("name").
		Values(city.Name).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		err = errors.Join(ErrInvalidQuery, err)
		return
	}

	err = pgxscan.Get(ctx, repo.pool, &id, query, args...)
	return
}

func (repo *city) List(ctx context.Context) (cities []models.City, err error) {
	query, args, err := repo.queryBuilder.
		Select("id", "name").
		From("city").
		OrderBy("id").
		ToSql()
	if err != nil {
		err = errors.Join(ErrInvalidQuery, err)
		return
	}

	err = pgxscan.Select(ctx, repo.pool, &cities, query, args...)
	return
}

func (repo *city) Update(ctx context.Context, city models.City) (err error) {
	query, args, err := repo.queryBuilder.
		Update("city").
		Set("name", city.Name).
		Where(squirrel.Eq{
			"id": city.ID,
		}).
		ToSql()
	if err != nil {
		err = errors.Join(ErrInvalidQuery, err)
		return
	}

	_, err = repo.pool.Exec(ctx, query, args...)
	return
}
