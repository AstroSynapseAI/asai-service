package rest

import (
	"net/http"

	"github.com/AstroSynapseAI/app-service/sdk/crud"
)

type CRUDController[T any] struct {
	Controller
	Repository crud.Repository[T]
}

func NewCRUDController[T any](repo crud.Repository[T]) *CRUDController[T] {
	return &CRUDController[T]{
		Repository: repo,
	}
}

func (ctrl *CRUDController[T]) Create(ctx *Context) {
	var record T

	err := ctx.JsonDecode(&record)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	newRecord, err := ctrl.Repository.Create(record)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusCreated, newRecord)
}

func (ctrl *CRUDController[T]) ReadAll(ctx *Context) {
	records, err := ctrl.Repository.ReadAll()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, records)
}

func (ctrl *CRUDController[T]) Read(ctx *Context) {
	record, err := ctrl.Repository.Read(ctx.GetID())
	if err != nil {
		ctx.JsonResponse(http.StatusNotFound, nil)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

func (ctrl *CRUDController[T]) Update(ctx *Context) {
	var record T

	err := ctx.JsonDecode(&record)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	updatedRecord, err := ctrl.Repository.Update(ctx.GetID(), record)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, updatedRecord)
}

func (ctrl *CRUDController[T]) Destroy(ctx *Context) {
	err := ctrl.Repository.Delete(ctx.GetID())
	if err != nil {
		ctx.Response.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Response.WriteHeader(http.StatusOK)
}