package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type PluginsController struct {
	rest.Controller
	Plugin *repositories.PluginsRepository
}

func NewPluginsController(db *database.Database) *PluginsController {
	return &PluginsController{
		Plugin: repositories.NewPluginsRepository(db),
	}
}

func (ctrl *PluginsController) Run() {
	ctrl.Post("/save/active", ctrl.SaveActivePlugin)
	ctrl.Post("/{id}/toggle/active", ctrl.ToggleActivePlugin)
}

func (ctrl *PluginsController) ReadAll(ctx *rest.Context) {
	fmt.Println("PluginsController.ReadAll")
	records, err := ctrl.Plugin.Repo.ReadAll()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *PluginsController) SaveActivePlugin(ctx *rest.Context) {
	fmt.Println("AvatarsController.SavePlugins")

	var plugin models.ActivePlugin

	err := ctx.JsonDecode(&plugin)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Plugin.SaveActivePlugin(plugin)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *PluginsController) ToggleActivePlugin(ctx *rest.Context) {
	fmt.Println("AvatarsController.SetActivePlugin")

	var input struct {
		AvatarID     uint `json:"avatar_id"`
		ActivePlugin bool `json:"activePlugin"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Plugin.ToggleActivePlugin(input.AvatarID, ctx.GetID(), input.ActivePlugin)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}
