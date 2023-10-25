package main

import (
	"api/config"
	"api/internal/entity"
	"api/internal/infra/database"
	"api/internal/infra/webserver/handlers"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	conf, _ := config.LoadConfig(".")
	println(conf.DBDriver)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
