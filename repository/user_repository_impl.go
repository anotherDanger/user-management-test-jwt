package repository

import (
	"context"
	"database/sql"
	"user-management-test/domain"
	"user-management-test/helper"

	"github.com/sirupsen/logrus"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Register(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	query := "insert into new_users(username, fullname, password) values(?, ?, ?)"
	_, err := sql.ExecContext(ctx, query, entity.Username, entity.FullName, entity.Password)
	if err != nil {
		helper.NewLoggerConfigure("user_repository.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}

	response := domain.Domain{
		Username: entity.Username,
		FullName: entity.FullName,
		Password: entity.Password,
	}

	return &response, nil
}
