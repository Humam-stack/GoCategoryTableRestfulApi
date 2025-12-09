package main

import (
	"belajar-golang-api/app"
	"belajar-golang-api/controller"
	"belajar-golang-api/exception"
	"belajar-golang-api/helpers"
	"belajar-golang-api/middleware"
	"belajar-golang-api/repository"
	"belajar-golang-api/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.ConnectDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Panic Handler untuk Error Handling
	router.PanicHandler = exception.PanicHandler
	server := http.Server{
		Addr:    ":8000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
