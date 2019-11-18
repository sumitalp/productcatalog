package main

import (
	"github.com/sumitalp/productcatalog/db"
	"github.com/sumitalp/productcatalog/handler"
	"github.com/sumitalp/productcatalog/repository"
	"github.com/sumitalp/productcatalog/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := repository.NewUserRepository(d)
	as := repository.NewProductRepository(d)
	h := handler.NewHandler(us, as)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
