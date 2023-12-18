package controllers

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type AvatarsController struct {
	rest.Controller
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

func (ctrl *AvatarsController) ReadAll(ctx *rest.Context) {
	fmt.Println("AvatarsController.ReadAll")
}

func (ctrl *AvatarsController) SaveAvatar(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveAvatar")
}

func (ctrl *AvatarsController) GetAgents(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetAgents")
}

func (ctrl *AvatarsController) GetTools(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetTools")
}

func (ctrl *AvatarsController) GetPlugins(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetPlugins")
}

func (ctrl *AvatarsController) GetDocuments(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetDocuments")
}
