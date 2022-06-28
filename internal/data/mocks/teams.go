package mocks

import (
	"github.com/rbraddev/shift-rota/internal/data"
)

var mockTeam = &data.Team{
	ID:      1,
	Name:    "Network",
	Version: 1,
}

type TeamModel struct{}

func (t TeamModel) Get(id int64) (*data.Team, error) {
	switch id {
	case 1:
		return mockTeam, nil
	default:
		return nil, data.ErrRecordNotFound
	}
}

func (t TeamModel) Insert(team *data.Team) error {
	return nil
}

func (t TeamModel) Delete(id int64) error {
	return nil
}

func (t TeamModel) GetAll(name string, filters data.Filters) ([]*data.Team, data.Metadata, error) {
	return nil, data.Metadata{}, nil
}

func (t TeamModel) Update(team *data.Team) error {
	return nil
}
