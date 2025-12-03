package controller

import (
	"belajar-golang-api/helpers"
	"belajar-golang-api/model/web"
	"belajar-golang-api/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}

}
func (controller CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helpers.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResp{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CategoryUpdateRequest := web.CategoryUpdateRequest{}
	helpers.ReadFromRequestBody(r, &CategoryUpdateRequest)

	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helpers.PanicIfError(err)

	CategoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), CategoryUpdateRequest)
	webResponse := web.WebResp{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helpers.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResp{
		Code:   200,
		Status: "OK",
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helpers.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResp{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}

func (controller CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryResponse := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResp{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}
