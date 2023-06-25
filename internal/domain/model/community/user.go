package community

type Role string

const (
	InvalidRole Role = "invalid"
	Admin       Role = "admin"
	User        Role = "user"
)

func (r Role) CanIssueInvitation() bool {
	return r == Admin
}
