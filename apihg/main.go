package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	dbConfig "github.com/lucas-sobierajski-oliveira/api/apihg/config/database"
	productController "github.com/lucas-sobierajski-oliveira/api/apihg/controllers/product"
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
