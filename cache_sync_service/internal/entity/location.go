package entity

import "fmt"

type Location struct {
	City  string
	State string
}

func (l Location) String() string {
	return fmt.Sprintf("%s - %s", l.City, l.State)
}
