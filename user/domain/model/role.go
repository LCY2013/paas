package model

type Role struct {
	ID         int64         `gorm:"primary_key;not_null;auto_increment"`
	RoleName   string        `gorm:"primary_key" json:"role_name"`
	RoleStatus string        `gorm:"primary_key" json:"role_status"`
	Permission []*Permission `gorm:"many2many:role_permission" json:"permission"`
}
