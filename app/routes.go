package app

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/AstroSynapseAI/app-service/controllers"
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	"github.com/gorilla/handlers"
)

type Routes struct {
	DB *database.Database
}

var _ rest.Routes = (*Routes)(nil)

func NewRoutes(db *database.Database) *Routes {
	return &Routes{
		DB: db,
	}
}

func (routes *Routes) LoadRoutes(router *rest.Rest) {
	//TMP API Controller
	// routes.rest.Route("/api").MapController(controllers.NewApiController(routes.DB)).Init()

	// CUSTOM CONTROLLERS
	// Route custom UsersController
	router.Route("/api/users").MapController(controllers.NewUsersController(routes.DB))
	// Route custom AvatarsController
	router.Route("/api/avatars").MapController(controllers.NewAvatarsController(routes.DB))
	// Route custom AgentsController
	router.Route("/api/agents").MapController(controllers.NewAgentsController(routes.DB))
	// Route custom DocumentsController
	router.Route("/api/documents").MapController(controllers.NewDocumentsController(routes.DB))

	// CRUD CONTROLLERS
	// Routing CRUD Tools controller
	toolsCtrl := rest.NewCRUDController[models.Tool](
		models.Tool{},
		gorm.NewRepository[models.Tool](routes.DB, models.Tool{}),
	)
	router.Route("/api/tools").MapController(toolsCtrl)

	// Routing CRUD Plugins controller
	pluginsCtr := rest.NewCRUDController[models.Plugin](
		models.Plugin{},
		gorm.NewRepository[models.Plugin](routes.DB, models.Plugin{}),
	)
	router.Route("/api/plugins").MapController(pluginsCtr)

	// Routing CRUD LLMs controller
	llmsCtrl := rest.NewCRUDController[models.LLM](
		models.LLM{},
		gorm.NewRepository[models.LLM](routes.DB, models.LLM{}),
	)
	router.Route("/api/llms").MapController(llmsCtrl)

	// Routing CRUD Roles controller
	rolesCtrl := rest.NewCRUDController[models.Role](
		models.Role{},
		gorm.NewRepository[models.Role](routes.DB, models.Role{}),
	)
	router.Route("/api/roles").MapController(rolesCtrl)

	// Routing CRUD Accounts controller
	accountsCtrl := rest.NewCRUDController[models.Account](
		models.Account{},
		gorm.NewRepository[models.Account](routes.DB, models.Account{}),
	)
	router.Route("/api/accounts").MapController(accountsCtrl)
}

func (routes *Routes) LoadMiddlewares(router *rest.Rest) {
	// TMP API Auth Middleware
	router.Mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				tokenString := r.Header.Get("Authorization")
				publicRoutes := []string{
					"/api/users/login",
					"/api/users/register",
					"/api/users/register/invite/",
				}

				for _, validRoute := range publicRoutes {
					if r.URL.Path == validRoute {
						next.ServeHTTP(w, r)
						return
					}
				}

				usersRepo := repositories.NewUsersRepository(routes.DB)
				_, err := usersRepo.GetByAPIToken(tokenString)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	})

	if os.Getenv("ENVIRONMENT") == "LOCAL DEV" {
		// CORS Middleware
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
		originsOk := handlers.AllowedOrigins([]string{"http://localhost:5173"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

		cors := handlers.CORS(headersOk, originsOk, methodsOk)
		router.Mux.Use(cors)

		// LOGGING MIDDLEWARE
		router.Mux.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasPrefix(r.URL.Path, "/api") {
					log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

					if r.Method == "POST" || r.Method == "PUT" {
						bodyBytes, err := io.ReadAll(r.Body)
						if err != nil {
							log.Printf("Error reading body: %v", err)
							http.Error(w, "can't read body", http.StatusBadRequest)
							return
						}

						// After reading the body, you need to replace it for further handlers
						r.Body = ioutil.NopCloser(strings.NewReader(string(bodyBytes)))

						log.Printf("Body: %s\n", string(bodyBytes))
					}
				}
				next.ServeHTTP(w, r)
			})
		})
	}

}
