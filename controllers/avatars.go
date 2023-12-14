package controllers

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type Avatars struct {
	rest.RestController
	Repo crud.Repository[models.Avatar]
}

func NewAvatarsController(db *database.Database) *Avatars {
	return &Avatars{
		Repo: gorm.NewRepository[models.Avatar](db, models.Avatar{}),
	}
}

func (ctrl *Avatars) Run() {

}

func (ctrl *Avatars) CreateAvatar(ctx *rest.Context) {}

func (ctrl *Avatars) GetAvatar(ctx *rest.Context) {}

func (ctrl *Avatars) UpdateAvatar(ctx *rest.Context) {}

func (ctrl *Avatars) GetAgents(ctx *rest.Context) {}

func (ctrl *Avatars) GetPlugins(ctx *rest.Context) {}

func (ctrl *Avatars) GetDocuments(ctx *rest.Context) {}
