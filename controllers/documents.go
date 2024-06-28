package controllers

import (
	"fmt"

	"github.com/AstroSynapseAI/asai-service/repositories"
	"github.com/AstroSynapseAI/asai-service/sdk/rest"
	"github.com/GoLangWebSDK/crud/database"
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

func (ctrl *DocumentsController) UploadDocument(ctx *rest.Context) {
	fmt.Println("DocumentsController.UploadDocument")
}
