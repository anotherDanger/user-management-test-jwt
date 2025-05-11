package service

import (
	"context"
	"database/sql"
	"user-management-test/domain"
	"user-management-test/helper"
	"user-management-test/repository"
	"user-management-test/web"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo repository.UserRepository
	db   *sql.DB
}

func NewUserService(repo repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		repo: repo,
		db:   db,
	}
}

func (svc *UserServiceImpl) Register(ctx context.Context, request *web.Request) (*web.Response, error) {
	tx, err := svc.db.Begin()
	if err != nil {
		helper.NewLoggerConfigure("user_service.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		helper.NewLoggerConfigure("user_service.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}
	entity := domain.Domain{
		Username: request.Username,
		FullName: request.FullName,
		Password: string(hashedPassword),
	}

	result, err := svc.repo.Register(ctx, tx, &entity)
	if err != nil {
		helper.NewLoggerConfigure("user_service.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	token, err := helper.NewGenerateJwt(request.Username)
	if err != nil {
		return nil, err
	}

	response := web.Response{
		Username: result.Username,
		FullName: result.FullName,
		Token:    token,
	}

	return &response, nil

}

func (svc *UserServiceImpl) Login(ctx context.Context, request *web.Request) (*web.Response, error) {
	tx, err := svc.db.Begin()
	if err != nil {
		helper.NewLoggerConfigure("user_service.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}
	entity := &domain.Domain{
		Username: request.Username,
	}
	user, err := svc.repo.Login(ctx, tx, entity)
	if err != nil {
		helper.NewLoggerConfigure("user_service.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	hashedPassword := user.Password
	inputPassword := request.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		helper.NewLoggerConfigure("user_service.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}

	token, err := helper.NewGenerateJwt(request.Username)
	if err != nil {
		helper.NewLoggerConfigure("user_service.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return nil, err
	}
	response := &web.Response{
		Username: request.Username,
		FullName: request.FullName,
		Token:    token,
	}

	return response, nil

}
