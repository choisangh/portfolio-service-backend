package main

import (
	"fmt"

	"github.com/choisangh/board-crud-backend/pkg/api"
	"github.com/choisangh/board-crud-backend/pkg/db"
	"github.com/choisangh/board-crud-backend/pkg/router"
)

func main() {
	dbHandler, err := db.NewAndConnectGorm("user:password@tcp(localhost:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error")
	}
	apis := api.NewAPI(dbHandler)
	r := router.Router(apis)

	r.Run(":5001")
}
