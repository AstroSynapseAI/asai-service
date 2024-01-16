package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type ToolsController struct {
	rest.Controller
	Tool *repositories.ToolsRepository
}

func NewToolsController(db *database.Database) *ToolsController {
	return &ToolsController{
		Tool: repositories.NewToolsRepository(db),
	}
}

func (ctrl *ToolsController) Run() {
	ctrl.Post("/save/active", ctrl.SaveActiveTool)
	ctrl.Post("/{id}/toggle/active", ctrl.ToggleActiveTool)
}

func (ctrl *ToolsController) ReadAll(ctx *rest.Context) {
	fmt.Println("ToolsController.ReadAll")
	records, err := ctrl.Tool.Repo.ReadAll()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *ToolsController) SaveActiveTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveTool")

	var tool models.ActiveTool

	err := ctx.JsonDecode(&tool)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Tool.SaveActiveTool(tool)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *ToolsController) ToggleActiveTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.ToggleActiveTool")

	var input struct {
		AvatarID   uint `json:"avatar_id"`
		ActiveTool bool `json:"active_tool"`
	}

	err := ctx.JsonDecode(&input)

	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Tool.ToggleActiveTool(input.AvatarID, ctx.GetID(), input.ActiveTool)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}
