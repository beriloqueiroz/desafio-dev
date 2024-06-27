package entity

type Notification struct {
	ID                   string
	User                 User
	ScheduleNotification ScheduleNotification
	Message              string
}

func NewNotification(id string, user User, scheduleNotification ScheduleNotification, message string) (*Notification, error) {
	return &Notification{
		ID:                   id,
		User:                 user,
		ScheduleNotification: scheduleNotification,
		Message:              message,
	}, nil
}
