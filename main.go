package main

import (
	"github.com/AlbatozK/go_backend_boilerplate/db"
	"github.com/AlbatozK/go_backend_boilerplate/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() (*gin.Engine, error) {
	r := gin.Default()
	apiGroup := r.Group("/api")
	{
		router.NewUserRouter().Init(apiGroup)
	}
	return r, nil
}

func main() {
	err := godotenv.Load(".env.production")
	if err != nil {
		panic(err)
	}
	db.Init()
	router, err := setupRouter()
	if err != nil {
		panic(err)
	}
	router.Run(":8000")
}
