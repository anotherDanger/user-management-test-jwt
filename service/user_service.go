package service

import (
	"context"
	"user-management-test/web"
)

type UserService interface {
	Register(ctx context.Context, request *web.Request) (*web.Response, error)
	Login(ctx context.Context, request *web.Request) (*web.Response, error)
}
