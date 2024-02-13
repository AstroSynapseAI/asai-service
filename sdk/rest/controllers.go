package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

// REST CONTROLLER
type RestController struct {
	Router *Rest
}

func NewController(router *Rest) *RestController {
	return &RestController{
		Router: router,
	}
}

func (ctrl *RestController) Get(path string, handler HandlerFunc) {
	ctrl.createHandler(path, handler).Methods("GET")
}

func (ctrl *RestController) Post(path string, handler HandlerFunc) {
	ctrl.createHandler(path, handler).Methods("POST")
}

func (ctrl *RestController) Put(path string, handler HandlerFunc) {
	ctrl.createHandler(path, handler).Methods("PUT")
}

func (ctrl *RestController) Delete(path string, handler HandlerFunc) {
	ctrl.createHandler(path, handler).Methods("DELETE")
}

func (ctrl *RestController) createHandler(path string, handler HandlerFunc) *mux.Route {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(NewContext(w, r))
	}
	return ctrl.Router.SubRouter.HandleFunc(path, h)
}

// MAPPED CONTROLLER
type MappedController struct {
	Router    *Rest
	SubRouter *mux.Router
	Route     string
	Handler   RestHandler
}

func NewMappedController(router *Rest) MappedController {
	return MappedController{
		Router:    router,
		SubRouter: router.SubRouter,
		Route:     router.currentRoute,
		Handler:   router.currentHandler,
	}
}

func (ctrl *MappedController) Map() {
	ctrl.Handler.SetRouter(ctrl.Router)
	ctrl.Handler.Run()

	// GET
	GetHandler := func(w http.ResponseWriter, r *http.Request) {
		ctrl.Handler.Read(NewContext(w, r))
	}
	ctrl.SubRouter.HandleFunc("/{id}", GetHandler).Methods("GET")

	// GET ALL
	GetAllHandler := func(w http.ResponseWriter, r *http.Request) {
		ctrl.Handler.ReadAll(NewContext(w, r))
	}
	ctrl.SubRouter.HandleFunc("", GetAllHandler).Methods("GET")

	// POST
	PostHandler := func(w http.ResponseWriter, r *http.Request) {
		ctrl.Handler.Create(NewContext(w, r))
	}
	ctrl.SubRouter.HandleFunc("", PostHandler).Methods("POST")

	// PUT
	PutHandler := func(w http.ResponseWriter, r *http.Request) {
		ctrl.Handler.Update(NewContext(w, r))
	}
	ctrl.SubRouter.HandleFunc("/{id}", PutHandler).Methods("PUT")

	// DELETE
	DeleteHandler := func(w http.ResponseWriter, r *http.Request) {
		ctrl.Handler.Destroy(NewContext(w, r))
	}
	ctrl.SubRouter.HandleFunc("/{id}", DeleteHandler).Methods("DELETE")
}

// (BASE) CONTROLLER
type Controller struct {
	RestController
}

var _ RestHandler = &Controller{}

func (ctrl *Controller) SetRouter(router *Rest) {
	ctrl.Router = router
}

func (ctrl *Controller) Run() {}

func (ctrl *Controller) Create(ctx *Context) {}

func (ctrl *Controller) Read(ctx *Context) {}

func (ctrl *Controller) ReadAll(ctx *Context) {}

func (ctrl *Controller) Update(ctx *Context) {}

func (ctrl *Controller) Destroy(ctx *Context) {}
