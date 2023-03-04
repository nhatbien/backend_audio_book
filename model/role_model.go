package model

/*=
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
	Id          uint64 `gorm:"primaryKey;autoIncrement"`
	RoleName    string `json:"role_name,omitempty"  gorm:"size:255;uniqueIndex"`
	Description string `json:"role_description,omitempty" `
}

type Permission struct {
	Id             uint   `json:"permission_id" db:"permission_id, omitempty"`
	PermissionName string `json:"permission_name,omitempty" db:"permission_name, omitempty" gorm:"uniqueIndex"`
}
