package services

import (
	"belajar-golang-api/exception"
	"belajar-golang-api/helpers"
	"belajar-golang-api/model/domain"
	"belajar-golang-api/model/web"
	"belajar-golang-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepo repository.CategoryRepository
	db           *sql.DB
	Validation   *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepo: categoryRepo,
		db:           db,
		Validation:   validate,
	}
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := c.Validation.Struct(request)
	helpers.PanicIfError(err)

	tx, err := c.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.PanicRollBack(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = c.CategoryRepo.Save(ctx, tx, category)

	return helpers.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := c.Validation.Struct(request)
	helpers.PanicIfError(err)

	tx, err := c.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.PanicRollBack(tx)

	category, err := c.CategoryRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name
	category = c.CategoryRepo.Save(ctx, tx, category)

	return helpers.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := c.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.PanicRollBack(tx)

	category, err := c.CategoryRepo.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	c.CategoryRepo.Delete(ctx, tx, category)

}

func (c *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.PanicRollBack(tx)

	category, err := c.CategoryRepo.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helpers.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := c.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.PanicRollBack(tx)

	categories := c.CategoryRepo.FindAll(ctx, tx)

	return helpers.ToCategoryResponses(categories)
}
