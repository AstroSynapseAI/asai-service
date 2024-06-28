package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/asai-service/models"
	"github.com/AstroSynapseAI/asai-service/repositories"
	"github.com/AstroSynapseAI/asai-service/sdk/rest"
	"github.com/GoLangWebSDK/crud/database"
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

func (ctrl *PluginsController) Read(ctx *rest.Context) {
	fmt.Println("PluginsController.Read")
	record, err := ctrl.Plugin.Repo.Read(ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *PluginsController) SaveActivePlugin(ctx *rest.Context) {
	fmt.Println("PluginsController.SaveActivePlugins")

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
	fmt.Println("PluginsController.ToggleActivePlugin")

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
