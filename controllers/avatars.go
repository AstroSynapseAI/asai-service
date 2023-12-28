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
	ctrl.Get("/{id}/agents", ctrl.GetAgents)
	ctrl.Get("/{id}/tools", ctrl.GetTools)
	ctrl.Get("/{id}/plugins", ctrl.GetAgents)
	ctrl.Get("/{id}/documents", ctrl.GetDocuments)
}

func (ctrl *AvatarsController) Read(ctx *rest.Context) {
	fmt.Println("AvatarsController.Read")

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
	var avatar models.Avatar
	avatar.Name = input.AvatarName
	avatar.Slug = createSlug(input.AvatarName)
	avatar.Primer = input.AvatarPrimer
	avatar.LLMID = sql.NullInt64{
		Int64: int64(input.AvatarLLMID),
		Valid: true,
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

func (ctrl *AvatarsController) GetAgents(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetAgents")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetActiveAgents(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetTools(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetTools")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetActiveTools(ID)

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *AvatarsController) GetPlugins(ctx *rest.Context) {
	fmt.Println("AvatarsController.GetPlugins")
	ID := ctx.GetID()
	records := ctrl.Avatar.GetActivePlugins(ID)

	ctx.JsonResponse(http.StatusOK, records)
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
