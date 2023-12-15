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
	ctrl.Post("/{user_id}/accounts", ctrl.SaveAccount)
	ctrl.Get("/{user_id}/accounts", ctrl.GetAccount)
	ctrl.Get("/{user_id}/avatars/{avatar_slug}", ctrl.GetAvatar)
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
		_ = ctx.JsonResponse(http.StatusUnauthorized, nil)
		return
	}

	_ = ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) SaveAccount(ctx *rest.Context) {

}

func (ctrl *UsersController) GetAccount(ctx *rest.Context) {

}

func (ctrl *UsersController) DeleteAccount(ctx *rest.Context) {

}

func (ctrl *UsersController) GetAvatar(ctx *rest.Context) {

}
