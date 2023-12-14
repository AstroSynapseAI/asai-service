package controllers

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type AgentsController struct {
	rest.RestController
	Repo crud.Repository[models.Agent]
}

func NewAgentsController(db *database.Database) *AgentsController {
	return &AgentsController{
		Repo: gorm.NewRepository[models.Agent](db, models.Agent{}),
	}
}

func (ctrl *AgentsController) Run() {

}
