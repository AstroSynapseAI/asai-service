package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	"github.com/GoLangWebSDK/crud/database"
)

type LLMSController struct {
	rest.Controller
	LLM *repositories.LLMSRepository
}

func NewLLMSController(db *database.Database) *LLMSController {
	return &LLMSController{
		LLM: repositories.NewLLMSRepository(db),
	}
}

func (ctrl *LLMSController) Run() {
	ctrl.Post("/save/active", ctrl.SaveActiveLLM)
	ctrl.Post("/{id}/toggle/active", ctrl.ToggleActiveLLM)
}

func (ctrl *LLMSController) ReadAll(ctx *rest.Context) {
	fmt.Println("LLMSController.ReadAll")
	records, err := ctrl.LLM.Repo.ReadAll()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *LLMSController) Read(ctx *rest.Context) {
	fmt.Println("LLMSController.Read")
	record, err := ctrl.LLM.Repo.Read(ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *LLMSController) SaveActiveLLM(ctx *rest.Context) {
	fmt.Println("LLMSController.SaveLLM")

	var llm models.ActiveLLM

	err := ctx.JsonDecode(&llm)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.LLM.SaveActiveLLM(llm)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}

func (ctrl *LLMSController) ToggleActiveLLM(ctx *rest.Context) {
	fmt.Println("LLMSController.ToggleActiveLLM")

	var input struct {
		AvatarID  uint `json:"avatar_id"`
		ActiveLLM bool `json:"is_active"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	err = ctrl.LLM.ToggleActiveLLM(input.AvatarID, ctx.GetID(), input.ActiveLLM)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetStatus(http.StatusOK)
}
