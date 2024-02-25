package main

import (
	"log"
	"toko1/db"
	"toko1/handler"
	"toko1/repository"
	"toko1/routers"
	"toko1/services"

	"os"
	_ "toko1/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Tag Service API
// @version 1.0
// @description A tag service API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	db, err := db.GetDatabase(dbUsername, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err.Error())
	}

	//#USER
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserController(userService)

	//#CATEGORY
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryServices(categoryRepository)
	categoryHandler := handler.NewCategoryController(categoryService)

	//#PRODUCT
	productRepository := repository.NewProductRepository(db)
	productService := services.NewProductServices(productRepository, categoryRepository)
	productHandler := handler.NewProductController(productService)

	//#CART
	cartRepository := repository.NewCartRepository(db)
	cartService := services.NewCartService(cartRepository, productRepository, userRepository, categoryRepository)
	cartHandler := handler.NewCartController(cartService)

	//#TRANSACTION
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository, productRepository, userRepository, categoryRepository)
	transactionHandler := handler.NewTransactionController(transactionService)


	 //ROUTER
	router := gin.Default()
	routers.UserRoute(router, userHandler)
	routers.CategoryRoutes(router, categoryHandler)
	routers.ProductRoutes(router, productHandler)
	routers.CartRoute(router, cartHandler)
	routers.TransactionRoute(router, transactionHandler)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + port)

	// add swagger
	
}