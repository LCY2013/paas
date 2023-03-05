package repository

import (
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/user/domain/model"
	"github.com/jinzhu/gorm"
)

// IUserRepository 创建需要实现的接口
type IUserRepository interface {
	// InitTable 初始化表
	InitTable() error
	// FindUserByID 根据ID查处找数据
	FindUserByID(int64) (*model.User, error)
	// CreateUser 创建一条 user 数据
	CreateUser(*model.User) (int64, error)
	// DeleteUserByID 根据ID删除一条 user 数据
	DeleteUserByID(int64) error
	// UpdateUser 修改更新数据
	UpdateUser(*model.User) error
	// FindAll 查找user所有数据
	FindAll() ([]model.User, error)

	// AddRole 分配角色
	AddRole(*model.User, []*model.Role) error
	// UpdateRole 更新用户角色
	UpdateRole(*model.User, []*model.Role) error
	// DeleteRole 删除用户角色
	DeleteRole(*model.User, []*model.Role) error

	// IsRight 判断用户是否有对应的权限
	IsRight(string, int64) bool
}

// NewUserRepository 创建userRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(
		&model.User{},
		&model.Role{},
		&model.Permission{}).Error
}

// AddRole 为以及存在的用户添加角色
func (u *UserRepository) AddRole(user *model.User, roles []*model.Role) error {
	return u.mysqlDb.Model(&user).Association("Role").Append(roles).Error
}

// UpdateRole 更新角色
func (u *UserRepository) UpdateRole(user *model.User, roles []*model.Role) error {
	return u.mysqlDb.Model(*user).Association("Role").Replace(roles).Error
}

// DeleteRole 删除用户角色
func (u *UserRepository) DeleteRole(user *model.User, roles []*model.Role) error {
	return u.mysqlDb.Model(&user).Association("Role").Delete(roles).Error
}

// IsRight 监测当前用户是否具备权限
func (u *UserRepository) IsRight(action string, userID int64) bool {
	permission := &model.Permission{}
	sql := "select p.id from user u, user_role ur, role r, role_permission rp, permission p where p.permission_action=? and p.id = rp.permission_id and rp.role_id = r.id AND ur.role_id = r.id and ur.user_id = u.id and u.id = ?"
	if err := u.mysqlDb.Raw(sql, action, userID).Scan(permission); err != nil {
		log.Error(err)
	}
	// 可以写其他判断逻辑
	return permission.ID > 0
}

// FindUserByID 根据ID查找User信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userID).Error
}

// CreateUser 创建User信息
func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

// DeleteUserByID 根据ID删除User信息
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

// UpdateUser 更新User信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(user).Error
}

// FindAll 获取结果集
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}
