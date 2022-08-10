package main

import (
	"fmt"
	"net/http"

	"go_studies/web-store-app/db"
	"go_studies/web-store-app/routes"

	_ "github.com/lib/pq"
)

func main() {
	db.PingDb()

	routes.Router()

	fmt.Println("Servidor ouvindo na http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
