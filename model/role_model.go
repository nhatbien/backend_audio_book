package model

/*
type Role int

const (
	MEMBER Role = iota
	ADMIN
)

func (r Role) String() string {
	return []string{"MEMBER", "ADMIN"}[r]
}
*/
type Role struct {
	Id          uint   `json:"-" db:"id, omitempty"`
	RoleName    string `json:"role_name,omitempty" db:"role_name, omitempty" gorm:"uniqueIndex"`
	Description string `json:"role_description,omitempty" db:"role_description, omitempty" gorm:"uniqueIndex"`
}

type Permission struct {
	Id             uint   `json:"permission_id" db:"permission_id, omitempty"`
	PermissionName string `json:"permission_name,omitempty" db:"permission_name, omitempty" gorm:"uniqueIndex"`
}
