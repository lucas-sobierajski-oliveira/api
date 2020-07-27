package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Product string          `json:"product"`
	Plans   map[string]plan `json:"cycle"`
}

type plan struct {
	Id          int     `json:"id"`
	PriceRenew  float32 `json:"price_renew"`
	PriceOrder  float32 `json:"price_order"`
	Month       int     `json:"month"`
	Description string  `json:"description"`
	IdProduct   int     `json:"id_product"`
}

func Index(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		products := make(map[string]product)

		productRows, _ := db.Query("select * from products")
		defer productRows.Close()

		for productRows.Next() {
			var p product
			productRows.Scan(&p.Id, &p.Name, &p.Product)

			plansRows, _ := db.Query("select * from plans where id_product = ?", p.Id)
			defer plansRows.Close()

			p.Plans = make(map[string]plan)
			for plansRows.Next() {
				var pl plan
				plansRows.Scan(&pl.Id, &pl.PriceRenew, &pl.PriceOrder, &pl.Month, &pl.Description, &pl.IdProduct)
				p.Plans[pl.Description] = pl
			}

			products[p.Product] = p
		}

		c.JSON(200, gin.H{"products": products})
	}
}

func Show(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products := make(map[string]product)

		productRows, _ := db.Query("select * from products where id = ?", c.Param("id"))
		defer productRows.Close()

		for productRows.Next() {
			var p product
			productRows.Scan(&p.Id, &p.Name, &p.Product)

			plansRows, _ := db.Query("select * from plans where id_product = ?", p.Id)
			defer plansRows.Close()

			p.Plans = make(map[string]plan)
			for plansRows.Next() {
				var pl plan
				plansRows.Scan(&pl.Id, &pl.PriceRenew, &pl.PriceOrder, &pl.Month, &pl.Description, &pl.IdProduct)
				p.Plans[pl.Description] = pl
			}

			products[p.Product] = p
		}

		c.JSON(http.StatusOK, gin.H{"products": products})
	}
}
