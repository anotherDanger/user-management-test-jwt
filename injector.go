//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"user-management-test/controller"
	"user-management-test/helper"
	"user-management-test/repository"
	"user-management-test/service"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var ServerSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
	helper.NewDb,
	NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)),
	NewServer,
)

func InitServer() (*http.Server, func(), error) {
	wire.Build(ServerSet)
	return nil, nil, nil
}
