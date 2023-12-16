package controllers

import (
	"net/http"

	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type UsersController struct {
	rest.RestController
	User *repositories.UsersRepository
}

func NewUsersController(db *database.Database) *UsersController {
	return &UsersController{
		User: repositories.NewUsersRepository(db),
	}
}

func (ctrl *UsersController) Run() {
	ctrl.Post("/login", ctrl.Login)
	ctrl.Get("/{id}/accounts/{account_id}", ctrl.GetAccount)
	ctrl.Get("/{id}/avatars", ctrl.GetAvatar)
}

func (ctrl *UsersController) Login(ctx *rest.Context) {
	username := ctx.GetParam("username")
	password := ctx.GetParam("password")
	loggedIn := ctrl.User.Login(username, password)

	if !loggedIn {
		_ = ctx.JsonResponse(http.StatusUnauthorized, nil)
		return
	}

	user := ctrl.User.GetByUsername(username)
	if user == nil {
		_ = ctx.JsonResponse(http.StatusNotFound, nil)
		return
	}

	_ = ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) GetAccount(ctx *rest.Context) {
	accountID := ctx.GetID("account_id")
	account := ctrl.User.GetUserAccount(ctx.GetID(), accountID)
	
	if account != nil {
		_ = ctx.JsonResponse(http.StatusOK, account)
	}

	_ = ctx.JsonResponse(http.StatusNotFound, nil)
}

func (ctrl *UsersController) GetAvatar(ctx *rest.Context) {
	avatar := ctrl.User.GetUserAvatar(ctx.GetID())
	
	if avatar != nil {
		_ = ctx.JsonResponse(http.StatusOK, avatar)
	}

	_ = ctx.JsonResponse(http.StatusNotFound, nil)
}
