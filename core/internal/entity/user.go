package entity

import (
	"errors"
	cacheentity "github.com/beriloqueiroz/desafio-back/cache_sync_service/pkg/entity"
	"github.com/google/uuid"
	"strings"
	"time"
)

type User struct {
	ID         string
	Active     bool
	Email      string
	Phone      string
	Location   cacheentity.Location
	CreateTime time.Time
}

func NewUser(id string, active bool, email string, phone string, location cacheentity.Location) (*User, error) {
	user := &User{
		ID:         id,
		Phone:      phone,
		Email:      email,
		CreateTime: time.Now(),
		Active:     active,
		Location:   location,
	}
	if user.Validate() != nil {
		return nil, user.Validate()
	}
	return user, nil
}

func (u *User) Activate() {
	u.Active = true
}

func (u *User) Deactivate() {
	u.Active = false
}

func (u *User) Validate() error {
	var msg []string
	if uuid.Validate(u.ID) != nil {
		msg = append(msg, "id is invalid")
	}
	if u.Phone == "" {
		msg = append(msg, "phone is required")
	}
	if u.Email == "" {
		msg = append(msg, "email is required")
	}
	if len(u.Location.State) != 2 {
		msg = append(msg, "location state is invalid")
	}
	if u.Location.City == "" {
		msg = append(msg, "location city is required")
	}
	if len(msg) > 0 {
		return errors.New(strings.Join(msg, "; "))
	}
	return nil
}
