package repository

import (
	"context"
	"database/sql"
	"user-management-test/domain"
)

type UserRepository interface {
	Register(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error)
	Login(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error)
}
