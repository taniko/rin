package community

import "time"

type (
	ID          string
	Name        string
	DisplayName string
)

type Community struct {
	ID          ID
	Name        Name
	DisplayName DisplayName
	Access      Access
	CreatedAt   time.Time
}
