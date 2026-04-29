package main

import (
	"goBackend/internal/products"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type applicaiton struct {
	config config
}

type config struct {
	addr string
	db   string
}

func (a *applicaiton) mount() *gin.Engine {
	db, err := gorm.Open(sqlserver.Open(a.config.db), &gorm.Config{})
	if err != nil {
		log.Panicf("Database can't contected.")
	}

	app := gin.Default()
	app.Use(cors.Default())
	apiGroup := app.Group("/api")

	productRepository := products.NewRepository(db)
	productServices := products.NewService(productRepository)
	productHandler := products.NewHandler(productServices)

	productGroup := apiGroup.Group("/products")
	productGroup.GET("", productHandler.GetAllProducts)
	productGroup.POST("", productHandler.CreateProduct)
	productGroup.DELETE("/:id",productHandler.DeleteProduct)

	return app

}

func (a *applicaiton) run(app *gin.Engine) error {
	err := app.Run(a.config.addr)
	if err != nil {
		return err
	}
	return nil
}
