package repository

import (
	"github.com/LCY2013/paas/user/domain/model"
	"github.com/jinzhu/gorm"
)

// IRoleRepository 创建需要实现的接口
type IRoleRepository interface {
	// FindRoleByID 根据ID查处找数据
	FindRoleByID(int64) (*model.Role, error)
	// CreateRole 创建一条 user 数据
	CreateRole(*model.Role) (int64, error)
	// DeleteRoleByID 根据ID删除一条 role 数据
	DeleteRoleByID(int64) error
	// UpdateRole 修改更新数据
	UpdateRole(*model.Role) error
	// FindAll 查找 role 所有数据
	FindAll() ([]model.Role, error)

	// FindAllRoleById 根据ID查找所有角色
	FindAllRoleById([]int64) ([]*model.Role, error)

	// AddPermission 添加角色权限
	AddPermission(*model.Role, []*model.Permission) error
	// UpdatePermission 更新角色权限
	UpdatePermission(*model.Role, []*model.Permission) error
	// DeletePermission 删除角色权限
	DeletePermission(*model.Role, []*model.Permission) error
}

type RoleRepository struct {
	mysqlDb *gorm.DB
}

func (r RoleRepository) FindRoleByID(roleID int64) (*model.Role, error) {
	role := &model.Role{}
	return role, r.mysqlDb.Preload("Permission").First(role, roleID).Error
}

func (r RoleRepository) CreateRole(role *model.Role) (int64, error) {
	return role.ID, r.mysqlDb.Create(role).Error
}

func (r RoleRepository) DeleteRoleByID(roleID int64) error {
	role, err := r.FindRoleByID(roleID)
	if err != nil {
		return err
	}
	if err = r.mysqlDb.Model(&role).Association("Permission").Delete(role.Permission).Error; err != nil {
		return err
	}

	return nil
}

func (r RoleRepository) UpdateRole(role *model.Role) error {
	return r.mysqlDb.Preload("Permission").Updates(role).Error
}

func (r RoleRepository) FindAll() (allRole []model.Role, err error) {
	return allRole, r.mysqlDb.Find(allRole).Error
}

// FindAllRoleById 根据ID获取所有角色
func (r RoleRepository) FindAllRoleById(id []int64) (roleAll []*model.Role, err error) {
	return roleAll, r.mysqlDb.Find(&roleAll, id).Error
}

// AddPermission 为角色添加权限
func (r RoleRepository) AddPermission(role *model.Role, permissions []*model.Permission) error {
	return r.mysqlDb.Model(&role).Association("Permission").Append(permissions).Error
}

// UpdatePermission 为角色添加权限
func (r RoleRepository) UpdatePermission(role *model.Role, permissions []*model.Permission) error {
	return r.mysqlDb.Model(&role).Association("Permission").Replace(permissions).Error
}

// DeletePermission 删除角色权限
func (r RoleRepository) DeletePermission(role *model.Role, permissions []*model.Permission) error {
	return r.mysqlDb.Model(&role).Association("Permission").Delete(permissions).Error
}

// NewRoleRepository 创建RoleRepository
func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &RoleRepository{
		mysqlDb: db,
	}
}
