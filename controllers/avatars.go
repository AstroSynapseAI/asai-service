package controllers

import (
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
	DB     *database.Database
	Avatar *repositories.AvatarsRepository
	Agent  *repositories.AgentsRepository
	Plugin *repositories.PluginsRepository
	Tool   *repositories.ToolsRepository
	LLM    *repositories.LLMSRepository
	Doc    *repositories.DocumentsRepository
}

func NewAvatarsController(db *database.Database) *AvatarsController {
	return &AvatarsController{
		DB:     db,
		Avatar: repositories.NewAvatarsRepository(db),
		Agent:  repositories.NewAgentsRepository(db),
		Plugin: repositories.NewPluginsRepository(db),
		Tool:   repositories.NewToolsRepository(db),
		LLM:    repositories.NewLLMSRepository(db),
		Doc:    repositories.NewDocumentsRepository(db),
	}
}

func (ctrl *AvatarsController) Run() {
	ctrl.Post("/save", ctrl.SaveAvatar)

	ctrl.Get("/{id}/agents", ctrl.GetAgents)
	ctrl.Get("/{id}/agents/{agent_id}", ctrl.GetAgent)

	ctrl.Get("/{id}/plugins", ctrl.GetPlugins)
	ctrl.Get("/{id}/plugins/{plugin_id}", ctrl.GetPlugin)

	ctrl.Get("/{id}/tools", ctrl.GetTools)
	ctrl.Get("/{id}/tools/{tool_id}", ctrl.GetTool)

	ctrl.Get("/{id}/llms", ctrl.GetLLMS)
	ctrl.Get("/{id}/llms/{llm_id}", ctrl.GetLLM)

	ctrl.Get("/{id}/documents", ctrl.GetDocuments)
	ctrl.Get("/{id}/documents/{document_id}", ctrl.GetDocument)

	ctrl.Get("/{id}/session/{session_id}", ctrl.GetSession)
}

// AVATARS
func (ctrl *AvatarsController) Read(ctx *rest.Context) {
	fmt.Println("AvatarsController.Read")
	record, err := ctrl.Avatar.Repo.Read(ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *AvatarsController) SaveAvatar(ctx *rest.Context) {
	fmt.Println("AvatarsController.SaveAvatar")

	var input struct {
		AvatarID     uint   `json:"avatar_id,omitempty"`
		UserID       uint   `json:"user_id"`
		AvatarLLMID  uint   `json:"avatar_llm_id"`
		AvatarName   string `json:"avatar_name"`
		AvatarPrimer string `json:"avatar_primer"`
		IsPublic     bool   `json:"is_public"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		fmt.Println("Bad Request:", err)
		ctx.SetStatus(http.StatusBadRequest)
		return
	}
	avatar := models.Avatar{
		Name:     input.AvatarName,
		Slug:     createSlug(input.AvatarName),
		Primer:   input.AvatarPrimer,
		IsPublic: input.IsPublic,
		LLMID:    input.AvatarLLMID,
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

	avatar.DefaultPrimer = input.AvatarPrimer
	record, err := ctrl.Avatar.Create(input.UserID, avatar)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *AvatarsController) GetSession(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetSession")
	sessionID := ctx.GetParam("session_id")

	var history *models.ChatHistory

	err := ctrl.DB.Adapter.Gorm().Where(models.ChatHistory{SessionID: sessionID}).Find(&history).Error
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, history)
}

// ACTIVE AGENTS
func (ctrl *AvatarsController) GetAgents(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetAgents")
	ID := ctx.GetID()
	records := ctrl.Agent.GetActiveAgents(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetAgent(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetAgent")
	ID := ctx.GetID()
	agentID := ctx.GetID("agent_id")

	record, err := ctrl.Agent.GetActiveAgent(ID, agentID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

// ACTIVE PLUGINS
func (ctrl *AvatarsController) GetPlugins(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetPlugins")
	ID := ctx.GetID()
	records := ctrl.Plugin.GetActivePlugins(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetPlugin(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetPlugin")
	ID := ctx.GetID()
	pluginID := ctx.GetID("plugin_id")

	record, err := ctrl.Plugin.GetActivePlugin(ID, pluginID)

	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

// ACTIVE TOOLS
func (ctrl *AvatarsController) GetTools(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetTools")
	ID := ctx.GetID()
	records := ctrl.Tool.GetAvatarTools(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetTool(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetTool")
	ID := ctx.GetID()
	toolID := ctx.GetID("tool_id")

	record, err := ctrl.Tool.GetAvatarTool(toolID, ID)

	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

// ACTIVE LLMS
func (ctrl *AvatarsController) GetLLMS(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetLLMs")
	ID := ctx.GetID()
	records := ctrl.LLM.GetActiveLLMs(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetLLM(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetLLM")
	ID := ctx.GetID()
	llmID := ctx.GetID("llm_id")

	record, err := ctrl.LLM.GetActiveLLM(ID, llmID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

// DOCUMENTS
func (ctrl *AvatarsController) GetDocuments(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetDocuments")
	ID := ctx.GetID()
	records := ctrl.Doc.GetDocuments(ID)

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
