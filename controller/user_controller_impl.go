package controller

import (
	"encoding/json"
	"net/http"
	"user-management-test/helper"
	"user-management-test/service"
	"user-management-test/web"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type UserControllerImpl struct {
	svc service.UserService
}

func NewUserController(svc service.UserService) UserController {
	return &UserControllerImpl{
		svc: svc,
	}
}

func (ctrl *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if r.Body == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(helper.NewResponse(http.StatusBadRequest, "error", web.Response{}))
		helper.NewLoggerConfigure("controller.log", logrus.ErrorLevel, http.ErrBodyNotAllowed.Error(), logrus.ErrorLevel)
		return
	}

	var reqBody web.Request
	json.NewDecoder(r.Body).Decode(&reqBody)

	if reqBody.Username == "" || reqBody.FullName == "" || reqBody.Password == "" {
		w.WriteHeader(400)
		helper.NewLoggerConfigure("controller.log", logrus.InfoLevel, http.ErrBodyNotAllowed.Error(), logrus.InfoLevel)
		return
	}

	response, err := ctrl.svc.Register(r.Context(), &reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(helper.NewResponse(http.StatusBadRequest, "error", response))
		helper.NewLoggerConfigure("controller.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(helper.NewResponse(http.StatusCreated, "OK", response))

}

func (ctrl *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var reqBody web.Request
	json.NewDecoder(r.Body).Decode(&reqBody)

	if reqBody.Username == "" || reqBody.Password == "" {
		w.WriteHeader(400)
		return
	}

	response, err := ctrl.svc.Login(r.Context(), &reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(helper.NewResponse(http.StatusBadRequest, "error", response))
		helper.NewLoggerConfigure("controller.log", logrus.ErrorLevel, err.Error(), logrus.ErrorLevel)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(helper.NewResponse(200, "OK", response))
}
