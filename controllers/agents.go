package controllers

import (
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type AgentsController struct {
	rest.RestController
	Agent *repositories.AgentsRepository
}

func NewAgentsController(db *database.Database) *AgentsController {
	return &AgentsController{
		Agent: repositories.NewAgentsRepository(db),
	}
}

func (ctrl *AgentsController) Run() {

}
