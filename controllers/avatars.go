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
	ctrl.Post("/agents/{agent_id}/active", ctrl.ToggleActiveAgent)
	ctrl.Get("/{id}/agents", ctrl.GetAgents)
	ctrl.Get("/{id}/agents/{agent_id}", ctrl.GetAgent)

	ctrl.Post("/{id}/plugins/save", ctrl.SavePlugin)
	ctrl.Post("/{id}/plugins/{plugin_id}/active", ctrl.ToggleActivePlugin)
	ctrl.Get("/{id}/plugins", ctrl.GetPlugins)
	ctrl.Get("/{id}/plugins/{plugin_id}", ctrl.GetPlugin)

	ctrl.Post("/{id}/tools/save", ctrl.SaveTool)
	ctrl.Post("/{id}/tools/{tool_id}/active", ctrl.ToggleActiveTool)
	ctrl.Get("/{id}/tools", ctrl.GetTools)
	ctrl.Get("/{id}/tools/{tool_id}", ctrl.GetTool)

	ctrl.Post("/{id}/llms/save", ctrl.SaveLLM)
	ctrl.Post("/{id}/llms/{llm_id}/active", ctrl.ToggleActiveLLM)
	ctrl.Get("/{id}/llms", ctrl.GetLLMS)
	ctrl.Get("/{id}/llms/{llm_id}", ctrl.GetLLM)

	ctrl.Post("/documents/upload", ctrl.UploadDocument)
	ctrl.Get("/{id}/documents", ctrl.GetDocuments)
	ctrl.Get("/{id}/documents/{document_id}", ctrl.GetDocument)

}

// AVATARS
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

// ACTIVE AGENTS
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

func (ctrl *AvatarsController) ToggleActiveAgent(ctx *rest.Context) {
	fmt.Println("AvatarsController.SetActiveAgent")

	var input struct {
		ActiveAgent bool `json:"activeAgent"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Avatar.ToggleActiveAgent(ctx.GetID(), ctx.GetID("agent_id"), input.ActiveAgent)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
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

// ACTIVE PLUGINS
func (ctrl *AvatarsController) SavePlugin(ctx *rest.Context) {
	fmt.Println("AvatarsController.SavePlugins")

	var plugin models.ActivePlugin

	err := ctx.JsonDecode(&plugin)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Avatar.SaveActivePlugin(plugin)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *AvatarsController) ToggleActivePlugin(ctx *rest.Context) {
	fmt.Println("AvatarsController.SetActivePlugin")

	var input struct {
		ActivePlugin bool `json:"activePlugin"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Avatar.ToggleActivePlugin(ctx.GetID(), ctx.GetID("plugin_id"), input.ActivePlugin)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *AvatarsController) GetPlugins(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetPlugins")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetActivePlugins(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetPlugin(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetPlugin")
	ID := ctx.GetID()
	pluginID := ctx.GetID("plugin_id")

	record, err := ctrl.Avatar.GetActivePlugin(ID, pluginID)

	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

// TOOLS
func (ctrl *AvatarsController) SaveTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveTool")

	var tool models.ActiveTool

	err := ctx.JsonDecode(&tool)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Avatar.SaveActiveTool(tool)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *AvatarsController) ToggleActiveTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.ToggleActiveTool")

	var input struct {
		ActiveTool bool `json:"activeTool"`
	}

	err := ctx.JsonDecode(&input)

	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}
}

func (ctrl *AvatarsController) GetTools(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetTools")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetActiveTools(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetTool")
	ID := ctx.GetID()
	toolID := ctx.GetID("tool_id")

	record, err := ctrl.Avatar.GetActiveTool(ID, toolID)

	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

// LLMS
func (ctrl *AvatarsController) SaveLLM(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveLLM")

	var llm models.ActiveLLM

	err := ctx.JsonDecode(&llm)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Avatar.SaveActiveLLM(llm)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *AvatarsController) ToggleActiveLLM(ctx *rest.Context) {
	fmt.Println("AvatarsController.ToggleActiveLLM")

	var input struct {
		ActiveLLM bool `json:"activeLLM"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.Avatar.ToggleActiveLLM(ctx.GetID(), ctx.GetID("llm_id"), input.ActiveLLM)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *AvatarsController) GetLLMS(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetLLMs")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetActiveLLMs(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetLLM(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetLLM")
	ID := ctx.GetID()
	llmID := ctx.GetID("llm_id")

	record, err := ctrl.Avatar.GetActiveLLM(ID, llmID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

// DOCUMENTS
func (ctrl *AvatarsController) UploadDocument(ctx *rest.Context) {
	fmt.Println("AvatarsController.UploadDocument")

}

func (ctrl *AvatarsController) GetDocuments(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetDocuments")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetDocuments(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetDocument(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetDocument")
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
