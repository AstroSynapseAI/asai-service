package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type UsersController struct {
	rest.Controller
	User *repositories.UsersRepository
}

func NewUsersController(db *database.Database) *UsersController {
	return &UsersController{
		User: repositories.NewUsersRepository(db),
	}
}

func (ctrl *UsersController) Run() {
	ctrl.Post("/login", ctrl.Login)
	ctrl.Post("/register", ctrl.Register)
	ctrl.Post("/register/invite", ctrl.RegisterInvite)
	ctrl.Get("/invite", ctrl.CreateInvite)
	ctrl.Get("/{id}/accounts", ctrl.GetAccounts)
	ctrl.Get("/{id}/accounts/{account_id}", ctrl.GetAccount)
	ctrl.Get("/{id}/avatars", ctrl.GetAvatar)
}

// Default CRUD routes
func (ctrl *UsersController) ReadAll(ctx *rest.Context) {
	users, err := ctrl.User.Repo.ReadAll()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}
	ctx.JsonResponse(http.StatusOK, users)
}

func (ctrl *UsersController) Read(ctx *rest.Context) {
	fmt.Println("Fetching user")
	userID := ctx.GetID("id")
	user, err := ctrl.User.Repo.Read(userID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}
	ctx.JsonResponse(http.StatusOK, user)
}

// Custom routes
//
// create user invite
func (ctrl *UsersController) CreateInvite(ctx *rest.Context) {
	err := ctrl.User.CreateInvite()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

// register invited user
func (ctrl *UsersController) RegisterInvite(ctx *rest.Context) {
	var reqData struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		InviteToken string `json:"invite_token"`
	}

	fmt.Println("Register invite token: ")

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.ConfirmInvite(
		reqData.Username,
		reqData.Password,
		reqData.InviteToken,
	)

	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) Register(ctx *rest.Context) {
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.Register(reqData.Username, reqData.Password)
	if err != nil {
		ctx.SetStatus(http.StatusUnauthorized)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) Login(ctx *rest.Context) {
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.Login(reqData.Username, reqData.Password)
	if err != nil {
		ctx.SetStatus(http.StatusUnauthorized)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) GetAccounts(ctx *rest.Context) {
	userID := ctx.GetID("id")
	user, err := ctrl.User.Repo.Read(userID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}
	accounts := user.Accounts

	if accounts != nil {
		ctx.JsonResponse(http.StatusOK, accounts)
	}

	ctx.SetStatus(http.StatusNotFound)
}

func (ctrl *UsersController) GetAccount(ctx *rest.Context) {
	accountID := ctx.GetID("account_id")
	account, err := ctrl.User.GetUserAccount(ctx.GetID(), accountID)
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, account)
}

func (ctrl *UsersController) GetAvatar(ctx *rest.Context) {
	avatar, err := ctrl.User.GetUserAvatar(ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, avatar)
}
