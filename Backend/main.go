package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Search struct {
	Query string `form:"query"`
}

func main() {
  router := gin.Default()
  
  router.Use(cors.Default())
  
  router.GET("/", searchPage)
  
  router.Run(":8080") 
}

func searchPage(ctx *gin.Context) {
	var search Search

	err := ctx.ShouldBind(&search)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": search.Query})
	// ctx.String(http.StatusOK, search.Query)
}