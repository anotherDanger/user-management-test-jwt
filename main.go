package main

import (
	"net/http"
	"user-management-test/controller"
	"user-management-test/helper"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func NewRouter(ctrl controller.UserController) *httprouter.Router {
	r := httprouter.New()
	r.POST("/v1/register", ctrl.Register)
	r.POST("/v1/login", ctrl.Login)

	return r
}

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
}

func main() {
	server, cleanup, err := InitServer()
	if err != nil {
		helper.NewLoggerConfigure("server.log", logrus.FatalLevel, err.Error(), logrus.FatalLevel)
		panic(err)
	}

	defer cleanup()

	err = server.ListenAndServe()
	if err != nil {
		helper.NewLoggerConfigure("server.log", logrus.FatalLevel, err.Error(), logrus.FatalLevel)
		panic(err)
	}
}
