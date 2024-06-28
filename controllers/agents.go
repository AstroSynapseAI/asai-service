package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/asai-service/models"
	"github.com/AstroSynapseAI/asai-service/repositories"
	"github.com/AstroSynapseAI/asai-service/sdk/rest"
	"github.com/GoLangWebSDK/crud/database"
)

type AgentsController struct {
	rest.Controller
	Agent *repositories.AgentsRepository
	Tool  *repositories.ToolsRepository
}

func NewAgentsController(db *database.Database) *AgentsController {
	return &AgentsController{
		Agent: repositories.NewAgentsRepository(db),
		Tool:  repositories.NewToolsRepository(db),
	}
}

func (ctrl *AgentsController) Run() {
	ctrl.Post("/save/active", ctrl.SaveActiveAgent)
	ctrl.Post("/{id}/toggle/active", ctrl.ToggleActiveAgent)

	ctrl.Get("/{id}/tools", ctrl.GetTools)
	ctrl.Get("/{id}/tool/{tool_id}", ctrl.GetTool)
}

func (ctrl *AgentsController) ReadAll(ctx *rest.Context) {
	fmt.Println("AgentsController.ReadAll")
	records, err := ctrl.Agent.Repo.ReadAll()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AgentsController) Read(ctx *rest.Context) {
	fmt.Println("AgentsController.Read")
	record, err := ctrl.Agent.Repo.Read(ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *AgentsController) SaveActiveAgent(ctx *rest.Context) {
	fmt.Println("AgentsController.SaveActiveAgent")

	var activeAgent models.ActiveAgent
	err := ctx.JsonDecode(&activeAgent)
	if err != nil {
		fmt.Println(err)
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	record, err := ctrl.Agent.SaveActiveAgent(activeAgent)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *AgentsController) ToggleActiveAgent(ctx *rest.Context) {
	fmt.Println("AgentsController.ToggleActiveAgent")

	var input struct {
		AvatarID    uint `json:"avatar_id"`
		ActiveAgent bool `json:"is_active"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Agent.ToggleActiveAgent(input.AvatarID, ctx.GetID(), input.ActiveAgent)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *AgentsController) GetTools(ctx *rest.Context) {
	fmt.Println("AgentsController.GetTools")
	records := ctrl.Tool.GetAgentTools(ctx.GetID())
	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AgentsController) GetTool(ctx *rest.Context) {
	fmt.Println("AgentsController.GetTool")
	record, err := ctrl.Tool.GetAgentTool(ctx.GetID("agent_id"), ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}
	ctx.JsonResponse(http.StatusOK, record)
}
