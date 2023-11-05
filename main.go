package main

import (
	"github.com/aldhipradana/apidemo/config"
	_ "github.com/aldhipradana/apidemo/docs"
	"github.com/aldhipradana/apidemo/modules/user"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Ge     *gin.Engine
}

// @title API Demo Swagger API
// @version 1.0
// @description This is a sample server for API Demo.
// @host localhost:8082
// @BasePath /
func main() {
	router := gin.Default()
	c := config.GetConfig()
	app := App{
		Config: c,
		DB:     config.GetDb(),
		Ge:     router,
	}
	// Handle the Swagger route
	router.GET("/apiDocs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add your other routes here
	// router.GET("/users/:id", getUserByID)

	app.RegisterRoutes()

	router.Run(":8082")
}

func (a *App) RegisterRoutes() {
	// Register routes here
	user.NewHandler(a.DB).RegisterRoutes(a.Ge)
}
