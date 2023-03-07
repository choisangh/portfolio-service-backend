package api

import "github.com/choisangh/board-crud-backend/pkg/db"

type APIs struct {
	db *db.DBHandler
}

type Response struct {
	Res string `json:"res"`
}

func NewAPI(handler *db.DBHandler) *APIs {
	a := APIs{db: handler}

	return &a
}
