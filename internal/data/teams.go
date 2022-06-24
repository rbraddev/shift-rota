package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/rbraddev/shift-rota/internal/database"
	"github.com/rbraddev/shift-rota/internal/validator"
)

type Team struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Version int32  `json:"version"`
}

func ValidateTeam(v *validator.Validator, team *Team) {
	v.CheckField(validator.NotBlank(team.Name), "name", "must be provided")
}

type TeamModel struct {
	DB *database.DB
}

func (t TeamModel) Insert(team *Team) error {
	query := `
        INSERT INTO teams (name) 
        VALUES ($1)
        RETURNING id, version`

	args := []any{team.Name}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return t.DB.QueryRowContext(ctx, query, args...).Scan(&team.ID, &team.Version)
}

func (t TeamModel) Get(id int64) (*Team, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT id, name, version
        FROM teams
        WHERE id = $1`

	var team Team

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, id).Scan(
		&team.ID,
		&team.Name,
		&team.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &team, nil
}

func (t TeamModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM teams
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := t.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (t TeamModel) GetAll(name string, filters Filters) ([]*Team, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), id, name, version
        FROM teams
        WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')     
        ORDER BY %s %s, id ASC
        LIMIT $2 OFFSET $3`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{name, filters.limit(), filters.offset()}

	rows, err := t.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	teams := []*Team{}

	for rows.Next() {
		var team Team

		err := rows.Scan(
			&totalRecords,
			&team.ID,
			&team.Name,
			&team.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		teams = append(teams, &team)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return teams, metadata, nil
}

func (t TeamModel) Update(team *Team) error {
	query := `
        UPDATE teams
        SET name = $1, version = version + 1
        WHERE id = $2 AND version = $3
        RETURNING version`

	args := []any{
		team.Name,
		team.ID,
		team.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, args...).Scan(&team.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}
