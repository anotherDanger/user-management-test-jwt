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

func (repo *UserRepositoryImpl) Login(ctx context.Context, sql *sql.Tx, entity *domain.Domain) (*domain.Domain, error) {
	var user domain.Domain
	query := "select username, password from new_users where username = ?"
	err := sql.QueryRowContext(ctx, query, entity.Username).Scan(&user.Username, &user.Password)
	if err != nil {
		helper.NewLoggerConfigure("user_repository.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}

	return &user, nil

}
