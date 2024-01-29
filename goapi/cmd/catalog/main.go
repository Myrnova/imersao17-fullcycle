package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/devfullcycle/imersao17/goapi/internal/database"
	"github.com/devfullcycle/imersao17/goapi/internal/service"
	"github.com/devfullcycle/imersao17/goapi/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryServiceDB := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productServiceDB := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryServiceDB)
	webProductHandler := webserver.NewWebProductHandler(productServiceDB)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}