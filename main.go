package main

import (
	"GolangRestfulApi/app"
	"GolangRestfulApi/controller"
	"GolangRestfulApi/helper"
	"GolangRestfulApi/repository"
	"GolangRestfulApi/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//db, err := sql.Open("mysql", "5zgepw8ufzj9uzuj0f4i:pscale_pw_sz5r4lMlZXx4Qoo7fYRNeKprpWdnWspOMRixWufJk55@tcp(us-east.connect.psdb.cloud)/golang?tls=true")

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindAll)
	router.POST("/api/categories", categoryController.FindAll)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
