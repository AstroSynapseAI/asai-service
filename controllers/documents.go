package controllers

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
)

type DocumentsController struct {
	rest.RestController
	Repo crud.Repository[models.Document]
}

func NewDocumentsController(db *database.Database) *DocumentsController {
	return &DocumentsController{
		Repo: gorm.NewRepository[models.Document](db, models.Document{}),
	}
}

func (ctrl *DocumentsController) Run() {
	ctrl.Post("/upload", ctrl.UploadDocument)
}

func (ctrl *DocumentsController) UploadDocument(ctx *rest.Context) {}
