package implements

import (
	"context"
	"database/sql"
	"errors"
	"github.com/beriloqueiroz/desafio-back/core/internal/entity"
	"time"
)

type PostgresScheduleRepository struct {
	Db *sql.DB
}

type Schedule struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"start_time"`
	Status    string    `json:"status"`
}

func (p *PostgresScheduleRepository) Save(ctx context.Context, scheduleNotification *entity.ScheduleNotification) error {
	var s Schedule
	err := p.Db.QueryRowContext(ctx, "SELECT id, start_time, status FROM schedules WHERE id = $1", scheduleNotification.ID).Scan(
		&s.ID, &s.StartTime, &s.Status)
	if errors.Is(err, sql.ErrNoRows) {
		// insert
		_, err = p.Db.ExecContext(ctx, "INSERT INTO schedules (id, start_time, status) VALUES ($1, $2, $3)",
			scheduleNotification.ID, scheduleNotification.StartTime, scheduleNotification.Status)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	// update
	_, err = p.Db.ExecContext(ctx, "UPDATE schedules SET start_time = $1, status = $2  WHERE id = $3;",
		scheduleNotification.StartTime, scheduleNotification.Status, scheduleNotification.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresScheduleRepository) Delete(ctx context.Context, id string) error {
	_, err := p.Db.ExecContext(ctx, "DELETE FROM schedules WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresScheduleRepository) FindFirstPendingBeforeDate(ctx context.Context, date time.Time) (*entity.ScheduleNotification, error) {
	var s Schedule
	err := p.Db.QueryRowContext(ctx, "SELECT id, start_time, status FROM schedules WHERE start_time <= $1 and status= $2",
		date.Format("2006-01-02T15:04:05"), "Pending").Scan(&s.ID, &s.StartTime, &s.Status)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return entity.NewScheduleNotification(s.ID, s.StartTime, entity.Status(s.Status))
}
