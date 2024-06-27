package entity

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Status string

const (
	Pending           Status = "Pending"
	Processing        Status = "Processing"
	Executed          Status = "Executed"
	ExecutedWithError Status = "ExecutedWithError"
)

type ScheduleNotification struct {
	ID        string
	StartTime time.Time
	Status    Status
}

func NewScheduleNotification(id string, startTime time.Time, status Status) (*ScheduleNotification, error) {
	schedule := &ScheduleNotification{
		ID:        id,
		StartTime: startTime,
		Status:    status,
	}
	err := schedule.Validate()
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *ScheduleNotification) MarkExecuted() {
	s.Status = Executed
}

func (s *ScheduleNotification) MarkExecutedWithError() {
	s.Status = ExecutedWithError
}

func (s *ScheduleNotification) MarkProcessing() {
	s.Status = Processing
}

func (s *ScheduleNotification) Validate() error {
	var msg []string
	if uuid.Validate(s.ID) != nil {
		msg = append(msg, "id is invalid")
	}
	if len(msg) > 0 {
		return errors.New(strings.Join(msg, "; "))
	}
	return nil
}
