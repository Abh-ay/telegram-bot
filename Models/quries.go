package models

import (
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	GetQuery *sqlx.Stmt `query:"get-query"`
}
