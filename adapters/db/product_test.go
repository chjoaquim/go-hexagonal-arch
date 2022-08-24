package db_test

import (
	"database/sql"
	"github.com/chjoaquim/go-hexagonal-arch/adapters/db"
	"github.com/chjoaquim/go-hexagonal-arch/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	insertProduct(Db)
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE products (
    			"id" string,
    			"name" string,
    			"price" float,
    			"status" string
            );`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func insertProduct(db *sql.DB) {
	insert := `insert into products values ("P1", "Product_1", 10, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Panic(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("P1")

	require.Nil(t, err)
	require.Equal(t, "Product_1", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 1"
	product.Price = 10.99

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetID(), productResult.GetID())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "enabled", product.GetStatus())

}
