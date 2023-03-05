package model

type User struct {
	ID         int64   `gorm:"primary_key;not_null;auto_increment"`
	UserName   string  `gorm:"unique;not_null" json:"user_name"`
	UserEmail  string  `gorm:"unique;not_null" json:"user_email"`
	IsAdmin    bool    `gorm:"is_admin" json:"is_admin"`
	UserPwd    string  `gorm:"user_pwd" json:"user_pwd"`
	UserStatus int32   `gorm:"user_status" json:"user_status"`
	Role       []*Role `gorm:"many2many:user_role"`
}
