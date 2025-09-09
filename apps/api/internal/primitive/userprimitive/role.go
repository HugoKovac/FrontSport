package userprimitive

// Roles defines the type for the "roles" enum field.
type Roles string

// Roles values.
const (
	RoleAdmin Roles = "admin"
	RoleUser  Roles = "user"
)

func (ro Roles) String() string {
	return string(ro)
}

func (Roles) Values() (roles []string) {
	for _, r := range []Roles{
		RoleAdmin,
		RoleUser,
	} {
		roles = append(roles, string(r))
	}

	return
}
