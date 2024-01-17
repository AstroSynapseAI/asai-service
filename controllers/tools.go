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
	ctrl.Post("/save/avatar", ctrl.SaveAvatarTool)
	ctrl.Post("/save/agent", ctrl.SaveAgentTool)
	ctrl.Post("/{id}/toggle/avatar", ctrl.ToggleAvatarTool)
	ctrl.Post("/{id}/toggle/avatar", ctrl.ToggleAgentTool)
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

func (ctrl *ToolsController) SaveAvatarTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveAvatarTool")

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

func (ctrl *ToolsController) SaveAgentTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveAgentTool")

	var tool models.AgentTool

	err := ctx.JsonDecode(&tool)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Tool.SaveAgentTool(tool)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *ToolsController) ToggleAvatarTool(ctx *rest.Context) {
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

	err = ctrl.Tool.ToggleAvatarTool(input.AvatarID, ctx.GetID(), input.ActiveTool)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *ToolsController) ToggleAgentTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.ToggleAgentTool")

	var input struct {
		AgentID    uint `json:"agent_id"`
		ActiveTool bool `json:"active_tool"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Tool.ToggleAgentTool(input.AgentID, ctx.GetID(), input.ActiveTool)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}
