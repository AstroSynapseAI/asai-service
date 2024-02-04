package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/app-service/models"
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
	ctrl.Post("/invite", ctrl.CreateInvite)
	ctrl.Get("/{id}/accounts", ctrl.GetAccounts)
	ctrl.Get("/{id}/accounts/{account_id}", ctrl.GetAccount)
	ctrl.Get("/{id}/avatars", ctrl.GetAvatar)
}

// Default CRUD routes
func (ctrl *UsersController) ReadAll(ctx *rest.Context) {
	fmt.Println("UsersController.ReadAll")
	users, err := ctrl.User.GetAll()
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
	fmt.Println("UsersController.CreateInvite")
	var input struct {
		Username string `json:"username"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	record, err := ctrl.User.CreateInvite(input.Username)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
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
	fmt.Println("UsersController.Register")
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
	fmt.Println("UsersController.Login")
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
	fmt.Println("UsersController.GetAccounts")
	userID := ctx.GetID()

	accounts, err := ctrl.User.GetAccounts(userID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, accounts)
}

func (ctrl *UsersController) GetAccount(ctx *rest.Context) {
	accountID := ctx.GetID("account_id")
	account, err := ctrl.User.GetAccount(ctx.GetID(), accountID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, account)
}

func (ctrl *UsersController) GetAvatar(ctx *rest.Context) {
	fmt.Println("UsersController.GetAvatar")
	avatar, err := ctrl.User.GetUserAvatar(ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, avatar)
}

func (ctrl *UsersController) Update(ctx *rest.Context) {
	fmt.Println("UsersController.Update")
	userID := ctx.GetID()

	var record models.User
	record.ID = userID

	err := ctx.JsonDecode(&record)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.Update(record)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}
