package main

import (
	"net/http"

	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Search struct {
	Query string `form:"query"`
}

type Page struct {
	Title string `json:"title"`
	Description string `json:"description"`
	URL string `json:"url"`
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

	pages, err := getPages(search.Query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": pages})
}

func getPages(query string) ([]Page, error) {
	db, err := sql.Open("sqlite", "file:oggle-db.db")
    if err != nil {
        return nil, err
    }
    defer db.Close()

	stmt := `SELECT title, description, url FROM pages WHERE title LIKE ? ORDER BY score DESC LIMIT 10`
	rows, err := db.Query(stmt, query+"%")
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Page
	for rows.Next() {
		var p Page
		if err := rows.Scan(&p.Title, &p.Description, &p.URL); err != nil {
			return nil, err
		}
		results = append(results, p)
	}

	return results, nil
}