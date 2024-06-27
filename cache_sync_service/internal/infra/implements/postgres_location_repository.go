package implements

import (
	"context"
	"database/sql"
	"errors"
	"github.com/beriloqueiroz/desafio-back/cache_sync_service/internal/entity"
)

type PostgresLocationRepository struct {
	Db *sql.DB
}

type Location struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func (p *PostgresLocationRepository) ListUniqueLocations(ctx context.Context, page, size int) ([]entity.Location, error) {
	limit := size
	offset := limit * (page - 1)
	rows, err := p.Db.QueryContext(ctx, "SELECT DISTINCT city,state FROM users WHERE active=$1 ORDER BY city LIMIT $2 OFFSET $3",
		true, limit, offset)
	defer rows.Close()
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	locations := make([]entity.Location, 0)
	for rows.Next() {
		var loc Location
		if err := rows.Scan(&loc.City, &loc.State); err != nil {
			return nil, err
		}
		location := entity.Location{
			City:  loc.City,
			State: loc.State,
		}
		locations = append(locations, location)
	}

	return locations, nil
}
