package app

import (
	"github.com/gin-gonic/gin"
	"github.com/williamliang705/bookstore-oauth-api/src/domain/access_token"
	"github.com/williamliang705/bookstore-oauth-api/src/http"
	"github.com/williamliang705/bookstore-oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
