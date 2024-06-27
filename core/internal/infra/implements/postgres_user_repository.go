package implements

import (
	"context"
	"database/sql"
	"errors"
	"github.com/beriloqueiroz/desafio-back/core/internal/entity"
	"time"
)

type PostgresUserRepository struct {
	Db *sql.DB
}

type User struct {
	ID      string    `json:"id"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Active  bool      `json:"active"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	Created time.Time `json:"created"`
}

func (p *PostgresUserRepository) Find(ctx context.Context, id string) (*entity.User, error) {
	var u User
	err := p.Db.QueryRowContext(ctx, "SELECT id, email, phone, active, city, state, created FROM users WHERE id = $1", id).Scan(
		&u.ID, &u.Email, &u.Phone, &u.Active, &u.City, &u.State, &u.Created)
	if err != nil {
		return nil, err
	}
	location := entity.Location{
		City:  u.City,
		State: u.State,
	}
	return entity.NewUser(id, u.Active, u.Email, u.Phone, location)
}

func (p *PostgresUserRepository) Save(ctx context.Context, user *entity.User) error {
	var u User
	err := p.Db.QueryRowContext(ctx, "SELECT id, email, phone, active, city, state, created FROM users WHERE id = $1", user.ID).Scan(
		&u.ID, &u.Email, &u.Phone, &u.Active, &u.City, &u.State, &u.Created)
	if errors.Is(err, sql.ErrNoRows) {
		// insert
		_, err = p.Db.ExecContext(ctx, "INSERT INTO users (id, email, phone, active, city, state, created) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			user.ID, user.Email, user.Phone, user.Active, user.Location.City, user.Location.State, user.CreateTime)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	// update
	_, err = p.Db.ExecContext(ctx, "UPDATE users SET email = $1, phone = $2, active = $3, city = $4, state = $5  WHERE id = $6;",
		user.Email, user.Phone, user.Active, user.Location.City, user.Location.State, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresUserRepository) ListActives(ctx context.Context, page, size int) ([]entity.User, error) {
	limit := size
	offset := limit * (page - 1)
	rows, err := p.Db.QueryContext(ctx, "SELECT id,email,phone,active,city,state,created FROM users WHERE active=$1 ORDER BY id LIMIT $2 OFFSET $3",
		true, limit, offset)
	defer rows.Close()
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	users := make([]entity.User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Phone, &user.Active, &user.City, &user.State, &user.Created); err != nil {
			return nil, err
		}
		location := entity.Location{
			City:  user.City,
			State: user.State,
		}
		u, err := entity.NewUser(user.ID, user.Active, user.Email, user.Phone, location)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}

	return users, nil
}
