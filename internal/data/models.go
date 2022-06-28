package data

import (
	"errors"

	"github.com/rbraddev/shift-rota/internal/database"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Teams TeamModelInterface
}

func NewModels(db *database.DB) Models {
	return Models{
		Teams: TeamModel{DB: db},
	}
}
