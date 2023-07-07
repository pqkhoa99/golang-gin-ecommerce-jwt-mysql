package main

import (
	"golang-jwttoken/config"
	"golang-jwttoken/controller"
	"golang-jwttoken/helper"
	"golang-jwttoken/model"
	"golang-jwttoken/repository"
	"golang-jwttoken/router"
	"golang-jwttoken/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load enviroment variables", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	// Init repository
	usersRepository := repository.NewUsersRepositoryImpl(db)
	productRepository := repository.NewProductRepositoryImpl(db)

	// Init service
	authenticationService := service.NewAuthenticationServiceImpl(usersRepository, validate)
	productService := service.NewProductServiceImpl(productRepository, validate)

	// Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(usersRepository)
	productController := controller.NewProductController(productService)

	routes := router.NewRouter(usersRepository, authenticationController, usersController, productController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
