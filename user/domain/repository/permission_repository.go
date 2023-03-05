package repository

import (
	"github.com/LCY2013/paas/user/domain/model"
	"github.com/jinzhu/gorm"
)

// IPermissionRepository 创建需要实现的接口
type IPermissionRepository interface {
	// FindPermissionByID 根据ID查处找数据
	FindPermissionByID(int64) (*model.Permission, error)
	// CreatePermission 创建一条 user 数据
	CreatePermission(*model.Permission) (int64, error)
	// DeletePermissionByID 根据ID删除一条 role 数据
	DeletePermissionByID(int64) error
	// UpdatePermission 修改更新数据
	UpdatePermission(*model.Permission) error
	// FindAll 查找 role 所有数据
	FindAll() ([]model.Permission, error)

	// FindAllPermissionById 根据ID查找所有权限
	FindAllPermissionById([]int64) ([]*model.Permission, error)
}

type PermissionRepository struct {
	mysqlDb *gorm.DB
}

func (p PermissionRepository) FindPermissionByID(permissionID int64) (*model.Permission, error) {
	permission := &model.Permission{}
	return permission, p.mysqlDb.First(permission, permissionID).Error
}

func (p PermissionRepository) CreatePermission(permission *model.Permission) (int64, error) {
	return permission.ID, p.mysqlDb.Create(permission).Error
}

func (p PermissionRepository) DeletePermissionByID(permissionID int64) error {
	return p.mysqlDb.Where("id = ?", permissionID).Delete(&model.Permission{}).Error
}

func (p PermissionRepository) UpdatePermission(permission *model.Permission) error {
	return p.mysqlDb.Updates(permission).Error
}

func (p PermissionRepository) FindAll() (allPermission []model.Permission, err error) {
	return allPermission, p.mysqlDb.Find(allPermission).Error
}

func (p PermissionRepository) FindAllPermissionById(id []int64) (permissionAll []*model.Permission, err error) {
	return permissionAll, p.mysqlDb.Find(&permissionAll, id).Error
}

// NewPermissionRepository 创建PermissioinRepository
func NewPermissionRepository(mysqlDb *gorm.DB) IPermissionRepository {
	return &PermissionRepository{
		mysqlDb: mysqlDb,
	}
}
