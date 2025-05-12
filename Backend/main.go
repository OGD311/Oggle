package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Search struct {
	Query string `uri:"query" binding:"required"`
}

func main() {
  router := gin.Default()

  router.GET("/", func(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{"msg": "home"})
  })
  
  router.GET("/:query", func(ctx *gin.Context)  {
	var search Search
	if err := ctx.ShouldBindUri(&search); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"msg" : err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"query": search.Query})
	
  })
  router.Run(":8080") 
}