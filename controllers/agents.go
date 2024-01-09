package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type AgentsController struct {
	rest.Controller
	Agent *repositories.AgentsRepository
}

func NewAgentsController(db *database.Database) *AgentsController {
	return &AgentsController{
		Agent: repositories.NewAgentsRepository(db),
	}
}

func (ctrl *AgentsController) Run() {

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
