package main

import (
	"github.com/sumitalp/productcatalog/db"
	"github.com/sumitalp/productcatalog/router"
)

func main() {
	r := router.New()

	d := db.New()
	db.AutoMigrate(d)

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
