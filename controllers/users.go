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
	ctrl.Post("/register/invite/{token}", ctrl.RegisterInvite)
	ctrl.Get("/{id}/accounts", ctrl.GetAccounts)
	ctrl.Get("/{id}/accounts/{account_id}", ctrl.GetAccount)
	ctrl.Get("/{id}/avatars", ctrl.GetAvatar)
}

// Default CRUD routes
func (ctrl *UsersController) ReadAll(ctx *rest.Context) {
	users, err := ctrl.User.Repo.ReadAll()
	if err != nil {
		_ = ctx.JsonResponse(http.StatusInternalServerError, nil)
		return
	}
	_ = ctx.JsonResponse(http.StatusOK, users)
}

func (ctrl *UsersController) Read(ctx *rest.Context) {
	userID := ctx.GetID("id")
	user, err := ctrl.User.Repo.Read(userID)
	if err != nil {
		_ = ctx.JsonResponse(http.StatusInternalServerError, nil)
		return
	}
	_ = ctx.JsonResponse(http.StatusOK, user)
}

// Custom routes
func (ctrl *UsersController) RegisterInvite(ctx *rest.Context) {
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(reqData)
	if err != nil {
		_ = ctx.JsonResponse(http.StatusBadRequest, nil)
		return
	}

	inviteToken := ctx.GetParam("token")
	if user := ctrl.User.GetByInviteToken(inviteToken); user != nil {
		_ = ctx.JsonResponse(http.StatusNotFound, nil)
		return
	}

	if ctrl.User.Register(reqData.Username, reqData.Password) {
		_ = ctx.JsonResponse(http.StatusOK, nil)
		return
	}

	_ = ctx.JsonResponse(http.StatusInternalServerError, nil)
}

func (ctrl *UsersController) Register(ctx *rest.Context) {
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(reqData)
	if err != nil {
		_ = ctx.JsonResponse(http.StatusBadRequest, nil)
		return
	}

	if ctrl.User.Register(reqData.Username, reqData.Password) {
		_ = ctx.JsonResponse(http.StatusOK, nil)
		return
	}

	_ = ctx.JsonResponse(http.StatusInternalServerError, nil)
}

func (ctrl *UsersController) Login(ctx *rest.Context) {
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(reqData)
	if err != nil {
		_ = ctx.JsonResponse(http.StatusBadRequest, nil)
		return
	}

	loggedIn := ctrl.User.Login(reqData.Username, reqData.Password)

	if !loggedIn {
		_ = ctx.JsonResponse(http.StatusUnauthorized, nil)
		return
	}

	user := ctrl.User.GetByUsername(reqData.Username)
	if user == nil {
		_ = ctx.JsonResponse(http.StatusNotFound, nil)
		return
	}

	_ = ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) GetAccounts(ctx *rest.Context) {
	userID := ctx.GetID("id")
	user, err := ctrl.User.Repo.Read(userID)
	if err != nil {
		_ = ctx.JsonResponse(http.StatusInternalServerError, nil)
		return
	}
	accounts := user.Accounts

	if accounts != nil {
		_ = ctx.JsonResponse(http.StatusOK, accounts)
	}

	_ = ctx.JsonResponse(http.StatusNotFound, nil)
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
