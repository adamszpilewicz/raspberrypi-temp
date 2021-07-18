package data

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFoundRecord = errors.New("record not found")
)

type Models struct {
	Temperatures TempModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Temperatures: TempModel{DB: db},
	}
}
