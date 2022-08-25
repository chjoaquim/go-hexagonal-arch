package main

import (
	"github.com/chjoaquim/go-hexagonal-arch/adapters/http/server"
	"github.com/chjoaquim/go-hexagonal-arch/adapters/listdb"
	"github.com/chjoaquim/go-hexagonal-arch/application"
)

func main() {

	listDB := listdb.NewListDb()
	service := application.NewProductService(listDB)

	server.RunService(&service)
}
