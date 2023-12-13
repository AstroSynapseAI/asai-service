package app

import (
	"github.com/AstroSynapseAI/app-service/controllers"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
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
	// repo := repositories.NewApiRepository(routes.DB)

	routes.rest.Route("/api").MapController(controllers.NewApiController(routes.DB)).Init()
}

func (routes *Routes) LoadMiddlewares() {

}
