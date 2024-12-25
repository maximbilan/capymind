package database

type Role string

const (
	Admin Role = "admin"
)

func IsAdmin(role *Role) bool {
	if role == nil {
		return false
	}
	return *role == Admin
}
