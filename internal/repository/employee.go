package repository

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yekuanyshev/aiplus/internal/models"
)

type employee struct {
	pool         *pgxpool.Pool
	queryBuilder squirrel.StatementBuilderType
}

func NewEmployee(pool *pgxpool.Pool) *employee {
	return &employee{
		pool:         pool,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (repo *employee) Create(ctx context.Context, employee models.Employee) (id int64, err error) {
	query, args, err := repo.queryBuilder.
		Insert("employee").
		Columns(
			"phone",
			"first_name",
			"last_name",
			"middle_name",
			"city_id",
		).
		Values(
			employee.Phone,
			employee.FirstName,
			employee.LastName,
			employee.MiddleName,
			employee.CityID,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		err = errors.Join(err, ErrInvalidQuery)
		return
	}

	err = pgxscan.Get(ctx, repo.pool, &id, query, args...)
	return
}

func (repo *employee) ByID(ctx context.Context, id int64) (employee models.Employee, err error) {
	query, args, err := repo.queryBuilder.
		Select(
			"id",
			"phone",
			"first_name",
			"last_name",
			"middle_name",
			"city_id",
		).
		From("employee").
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		ToSql()
	if err != nil {
		err = errors.Join(err, ErrInvalidQuery)
		return
	}

	err = pgxscan.Get(ctx, repo.pool, &employee, query, args...)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return
	}
	return
}
