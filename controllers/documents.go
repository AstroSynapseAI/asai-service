package controllers

import (
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type DocumentsController struct {
	rest.Controller
	Document *repositories.DocumentsRepository
}

func NewDocumentsController(db *database.Database) *DocumentsController {
	return &DocumentsController{
		Document: repositories.NewDocumentsRepository(db),
	}
}

func (ctrl *DocumentsController) Run() {
	ctrl.Post("/upload", ctrl.UploadDocument)
}

func (ctrl *DocumentsController) UploadDocument(ctx *rest.Context) {}
