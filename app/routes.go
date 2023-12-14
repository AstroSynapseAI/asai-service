package app

import (
	"fmt"
	"net/http"

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
	//TMP API Controller
	routes.rest.Route("/api").MapController(controllers.NewApiController(routes.DB)).Init()

	// CUSTOM CONTROLLERS
	// Route custom UsersController
	routes.rest.Route("/api/users").MapController(controllers.NewUsersController(routes.DB)).Init()
	// Route custom AgentsController
	routes.rest.Route("/api/agents").MapController(controllers.NewAgentsController(routes.DB)).Init()
	// Route custom DocumentsController
	routes.rest.Route("/api/documents").MapController(controllers.NewDocumentsController(routes.DB)).Init()

	// CRUD CONTROLLERS
	// Routing CRUD Tools controller
	toolsCtrl := rest.NewCRUDController[models.Tool](
		models.Tool{},
		gorm.NewRepository[models.Tool](routes.DB, models.Tool{}),
	)
	routes.rest.Route("/api/tools").MapController(toolsCtrl).Init()

	// Routing CRUD Plugins controller
	pluginsCtr := rest.NewCRUDController[models.Plugin](
		models.Plugin{},
		gorm.NewRepository[models.Plugin](routes.DB, models.Plugin{}),
	)
	routes.rest.Route("/api/plugins").MapController(pluginsCtr).Init()

	// Routing CRUD LLMs controller
	llmsCtrl := rest.NewCRUDController[models.LLM](
		models.LLM{},
		gorm.NewRepository[models.LLM](routes.DB, models.LLM{}),
	)
	routes.rest.Route("/api/llms").MapController(llmsCtrl).Init()

	// Routing CRUD Roles controller
	rolesCtrl := rest.NewCRUDController[models.Role](
		models.Role{},
		gorm.NewRepository[models.Role](routes.DB, models.Role{}),
	)
	routes.rest.Route("/api/roles").MapController(rolesCtrl).Init()
}

func (routes *Routes) LoadMiddlewares() {
	routes.rest.Mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("Middleware called")

			next.ServeHTTP(w, r)
		})
	})

}
