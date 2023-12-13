package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/app-service/engine"
	"github.com/AstroSynapseAI/app-service/engine/chains"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	"github.com/thanhpk/randstr"
)

type ApiController struct {
	rest.RestController
	Repo *repositories.ApiRepository
}

var _ rest.Controller = (*ApiController)(nil)

func NewApiController(db *database.Database) *ApiController {
	return &ApiController{
		Repo: repositories.NewApiRepository(db),
	}
}

func (ctrl *ApiController) Run() {

	// Get Chat History endpoint.
	ctrl.Get("/chat/history/{session_id}", func(ctx *rest.Context) {
		sessionId := ctx.GetParam("session_id")
		history := ctrl.Repo.GetChatHistory(sessionId)

		_ = ctx.JsonResponse(http.StatusOK, history.ChatHistory)
	})

	// Send Chat Message endpoint.
	ctrl.Post("/chat/msg", func(ctx *rest.Context) {

		// Parse the incoming http request
		var request struct {
			SessionId  string `json:"session_id"`
			UserPrompt string `json:"user_prompt"`
		}

		err := ctx.JsonDecode(&request)
		if err != nil {
			fmt.Println("Bad Request: %w", err)
			_ = ctx.JsonResponse(http.StatusBadRequest, err)
			return
		}

		// Initialize Asai Chain
		asaiConfig := engine.NewConfig(ctrl.Repo.DB)
		asaiChain, _ := chains.NewAsaiChain(asaiConfig)
		asaiChain.SetSessionID(request.SessionId)

		// Send user prompt to Asai Chain
		response, err := asaiChain.Prompt(context.Background(), request.UserPrompt)
		if err != nil {
			fmt.Println(err)
			_ = ctx.JsonResponse(http.StatusInternalServerError, err)
			return
		}

		// Send response to user
		var responseJson struct {
			Content string `json:"content"`
		}

		responseJson.Content = response

		_ = ctx.JsonResponse(http.StatusOK, responseJson)
	})

	// Create new User session endpoint.
	ctrl.Get("/users/session", func(ctx *rest.Context) {
		sessionID := randstr.String(16)

		var reponseJson struct {
			SessionId string `json:"session_id"`
		}

		reponseJson.SessionId = sessionID

		_ = ctx.JsonResponse(http.StatusOK, reponseJson)
	})
}
