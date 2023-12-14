package app

import (
	"github.com/AstroSynapseAI/app-service/controllers"
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type Routes struct {
	rest *rest.Rest
	DB   *database.Database
}

var _ rest.Routes = (*Routes)(nil)

func NewRoutes(router *rest.Rest, db *database.Database) *Routes {
	return &Routes{
		rest: router,
		DB:   db,
	}
}

func (routes *Routes) LoadRoutes() {
	routes.rest.Route("/api").MapController(controllers.NewApiController(routes.DB)).Init()

	routes.rest.Route("/api/users").MapController(controllers.NewUsersController(routes.DB)).Init()

	agentsCtrl := rest.NewCRUDController[models.Agent](
		models.Agent{},
		gorm.NewRepository[models.Agent](routes.DB, models.Agent{}),
	)
	routes.rest.Route("/api/agents").MapController(agentsCtrl).Init()

	toolsCtrl := rest.NewCRUDController[models.Tool](
		models.Tool{},
		gorm.NewRepository[models.Tool](routes.DB, models.Tool{}),
	)
	routes.rest.Route("/api/tools").MapController(toolsCtrl).Init()

	pluginsCtr := rest.NewCRUDController[models.Plugin](
		models.Plugin{},
		gorm.NewRepository[models.Plugin](routes.DB, models.Plugin{}),
	)
	routes.rest.Route("/api/plugins").MapController(pluginsCtr).Init()

	llmsCtrl := rest.NewCRUDController[models.LLM](
		models.LLM{},
		gorm.NewRepository[models.LLM](routes.DB, models.LLM{}),
	)
	routes.rest.Route("/api/llms").MapController(llmsCtrl).Init()
}

func (routes *Routes) LoadMiddlewares() {

}
