package main

import (
	"database/sql"
	"fmt"

	dbConfig "github.com/api/apihg/config/database"
	productController "github.com/api/apihg/controllers/product"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sql.Open(dbConfig.DRIVER, fmt.Sprintf("%s:%s@/cursogo", dbConfig.USER, dbConfig.PASSWORD))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	routes := gin.Default()

	routes.Use(cors.Default())

	routes.GET("/products", productController.Index(db))

	routes.GET("/products/:id", productController.Show(db))

	routes.Run()
}
