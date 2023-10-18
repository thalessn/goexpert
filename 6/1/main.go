package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products values (?, ?, ?);")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Id, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Id)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}
	var product Product
	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		panic(err)
	}
	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {
	// user:password@tcp(localhost)/database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Test", 16.0)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	// product.Price = 20
	// err = updateProduct(db, product)
	// if err != nil {
	// 	panic(err)
	// }

	productSaved, err := selectProduct(db, "0d5407b9-51ea-49f7-8509-d8f99f35a575")
	if err != nil {
		panic(err)
	}
	fmt.Printf("O produto possui nome: %s e preco de %.2f", productSaved.Name, productSaved.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("O produto é chamado de %s e custa R$%.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(db, "0d5407b9-51ea-49f7-8509-d8f99f35a575")
	if err != nil {
		panic(err)
	}
}
