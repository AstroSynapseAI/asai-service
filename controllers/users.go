package controllers

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type Users struct {
	rest.RestController
	Repo crud.Repository[models.User]
}

func NewUsersController(db *database.Database) *Users {
	return &Users{
		Repo: gorm.NewRepository[models.User](db, models.User{}),
	}
}

func (ctrl *Users) Run() {
	ctrl.Post("/login", ctrl.Login)
	ctrl.Post("/{user_id}/accounts", ctrl.SaveAccount)
	ctrl.Get("/{user_id}/accounts", ctrl.GetAccount)
	ctrl.Get("/{user_id}/avatars/{avatar_slug}", ctrl.GetAvatar)
}

func (ctrl *Users) Login(ctx *rest.Context) {

}

func (ctrl *Users) SaveAccount(ctx *rest.Context) {

}

func (ctrl *Users) GetAccount(ctx *rest.Context) {

}

func (ctrl *Users) DeleteAccount(ctx *rest.Context) {

}

func (ctrl *Users) GetAvatar(ctx *rest.Context) {

}
