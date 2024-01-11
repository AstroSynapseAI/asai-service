package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/AstroSynapseAI/app-service/models"
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
	ctrl.Post("/save", ctrl.SaveAvatar)

	ctrl.Post("/agents/save", ctrl.SaveAgent)
	ctrl.Post("/agents/{agent_id}/active", ctrl.SetActiveAgent)
	ctrl.Get("/{id}/agents", ctrl.GetAgents)
	ctrl.Get("/{id}/agents/{agent_id}", ctrl.GetAgent)

	ctrl.Post("/{id}/plugins/save", ctrl.SavePlugin)
	ctrl.Post("/{id}/plugins/{plugin_id}/active", ctrl.SetActivePlugin)
	ctrl.Get("/{id}/plugins", ctrl.GetPlugins)
	ctrl.Get("/{id}/plugins/{plugin_id}", ctrl.GetPlugin)

	ctrl.Get("/{id}/documents", ctrl.GetDocuments)
}

func (ctrl *AvatarsController) SaveAvatar(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveAvatar")

	var input struct {
		AvatarID     uint   `json:"avatar_id,omitempty"`
		UserID       uint   `json:"user_id"`
		AvatarLLMID  uint   `json:"avatar_llm_id"`
		AvatarName   string `json:"avatar_name"`
		AvatarPrimer string `json:"avatar_primer"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		fmt.Println("Bad Request:", err)
		ctx.SetStatus(http.StatusBadRequest)
		return
	}
	avatar := models.Avatar{
		Name:          input.AvatarName,
		Slug:          createSlug(input.AvatarName),
		Primer:        input.AvatarPrimer,
		DefaultPrimer: input.AvatarPrimer,
		IsPublic:      false,
		LLMID: sql.NullInt64{
			Int64: int64(input.AvatarLLMID),
			Valid: true,
		},
	}

	if input.AvatarID != 0 {
		record, err := ctrl.Avatar.Update(input.AvatarID, avatar)
		if err != nil {
			ctx.SetStatus(http.StatusInternalServerError)
			return
		}

		ctx.JsonResponse(http.StatusOK, record)
		return
	}

	record, err := ctrl.Avatar.Create(input.UserID, avatar)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *AvatarsController) SaveAgent(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveAgent")

	var activeAgent models.ActiveAgent
	err := ctx.JsonDecode(&activeAgent)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Avatar.SaveActiveAgent(activeAgent)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *AvatarsController) SetActiveAgent(ctx *rest.Context) {
	fmt.Println("AvatarsController.SetActiveAgent")
}

func (ctrl *AvatarsController) GetAgents(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetAgents")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetActiveAgents(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetAgent(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetAgent")
	ID := ctx.GetID()
	agentID := ctx.GetID("agent_id")

	record, err := ctrl.Avatar.GetActiveAgent(ID, agentID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *AvatarsController) SavePlugin(ctx *rest.Context) {
	fmt.Println("AvatarsController.SavePlugins")
}

func (ctrl *AvatarsController) SetActivePlugin(ctx *rest.Context) {
	fmt.Println("AvatarsController.SetActivePlugin")
}

func (ctrl *AvatarsController) GetPlugins(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetPlugins")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetActivePlugins(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetPlugin(ctx *rest.Context) {

}

func (ctrl *AvatarsController) GetDocuments(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetDocuments")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetDocuments(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func createSlug(input string) string {
	// Lowercase the string
	input = strings.ToLower(input)

	// Use the regexp package to remove any non-alphanumeric characters
	reg, _ := regexp.Compile("[^a-z0-9]+")
	input = reg.ReplaceAllString(input, "-")

	// Remove any trailing hyphens
	input = strings.Trim(input, "-")

	return input
}
