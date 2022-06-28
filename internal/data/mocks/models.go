package mocks

import (
	"github.com/rbraddev/shift-rota/internal/data"
)

func NewMockModels() data.Models {
	return data.Models{
		Teams: TeamModel{},
	}
}
