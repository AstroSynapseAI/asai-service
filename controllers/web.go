package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

const (
	DefualtHomepageDir = "./web/static"
	DefualtAdminDir    = "./admin/static"
)

type WebController struct {
	router      *rest.Rest
	homepageDir string
	adminDir    string
}

func NewWebController(router *rest.Rest) *WebController {
	return &WebController{
		router:      router,
		homepageDir: DefualtHomepageDir,
		adminDir:    DefualtAdminDir,
	}
}

func (ctrl *WebController) Run() {
	ctrl.RunAdmin()
	ctrl.RunHomepage()
	ctrl.vueFallback()
}

func (ctrl *WebController) RunHomepage() {
	static := http.FileServer(http.Dir(ctrl.homepageDir))
	assets := http.FileServer(http.Dir(ctrl.homepageDir + "/assets"))

	ctrl.router.Mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assets))
	ctrl.router.Mux.Handle("/", static)
}

func (ctrl *WebController) RunAdmin() {
	static := http.FileServer(http.Dir(ctrl.adminDir))
	assets := http.FileServer(http.Dir(ctrl.adminDir + "/assets"))

	ctrl.router.Mux.PathPrefix("/admin/assets/").Handler(http.StripPrefix("/admin/assets/", assets))
	ctrl.router.Mux.Handle("/admin/", static)
}

func (ctrl *WebController) vueFallback() {
	admin := http.FileServer(http.Dir(ctrl.adminDir))
	ctrl.router.Mux.PathPrefix("/admin/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(ctrl.adminDir + r.URL.Path)
		_, err := os.Stat(path)

		if os.IsNotExist(err) {
			http.ServeFile(w, r, ctrl.adminDir+"/index.html")
			return
		}

		// If request is not for a directory, serve with the static file server as normal
		admin.ServeHTTP(w, r)
	})

	homepage := http.FileServer(http.Dir(ctrl.homepageDir))
	// Fallback to index.html for Vue Router
	ctrl.router.Mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(ctrl.homepageDir + r.URL.Path)
		_, err := os.Stat(path)

		if os.IsNotExist(err) {
			http.ServeFile(w, r, ctrl.homepageDir+"/index.html")
			return
		}

		// If request is not for a directory, serve with the static file server as normal
		homepage.ServeHTTP(w, r)
		return
	})
}
