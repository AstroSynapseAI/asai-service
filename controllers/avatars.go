package controllers

import (
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type AvatarsController struct {
	rest.RestController
	Avatar *repositories.AvatarsRepository
}

func NewAvatarsController(db *database.Database) *AvatarsController {
	return &AvatarsController{
		Avatar: repositories.NewAvatarsRepository(db),
	}
}

func (ctrl *AvatarsController) Run() {
	ctrl.Post("/{avatar_slug}", ctrl.SaveAvatar)
	ctrl.Get("/{avatar_slug}/agents", ctrl.GetAgents)
	ctrl.Get("/{avatar_slug}/tools", ctrl.GetTools)
	ctrl.Get("/{avatar_slug}/plugins", ctrl.GetAgents)
	ctrl.Get("/{avatar_slug}/documents", ctrl.GetDocuments)
}

func (ctrl *AvatarsController) SaveAvatar(ctx *rest.Context) {}

func (ctrl *AvatarsController) GetAgents(ctx *rest.Context) {}

func (ctrl *AvatarsController) GetTools(ctx *rest.Context) {}

func (ctrl *AvatarsController) GetPlugins(ctx *rest.Context) {}

func (ctrl *AvatarsController) GetDocuments(ctx *rest.Context) {}
