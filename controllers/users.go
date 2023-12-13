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

}

func (ctrl *Users) Login(ctx *rest.Context) {

}
