package service

import (
	"github.com/LCY2013/paas/user/domain/model"
	"github.com/LCY2013/paas/user/domain/repository"
)

// IPermissionDataService 这里是接口类型
type IPermissionDataService interface {
	AddPermission(permission *model.Permission) (int64, error)
	DeletePermission(permissionID int64) error
	UpdatePermission(permission *model.Permission) error
	FindPermissionByID(permissionID int64) (*model.Permission, error)
	FindAllPermission() ([]model.Permission, error)

	// FindAllPermissionById 根据ID查询所有权限
	FindAllPermissionById(ids []int64) ([]*model.Permission, error)
}

// NewPermissionDataService
// 创建
// 注意：返回值 IPermissionDataService 接口类型
func NewPermissionDataService(permissionRepository repository.IPermissionRepository) IPermissionDataService {
	return &PermissionDataService{PermissionRepository: permissionRepository}
}

type PermissionDataService struct {
	//注意：这里是 IPermissionRepository 类型
	PermissionRepository repository.IPermissionRepository
}

func (p PermissionDataService) AddPermission(permission *model.Permission) (int64, error) {
	return p.PermissionRepository.CreatePermission(permission)
}

func (p PermissionDataService) DeletePermission(permissionID int64) error {
	return p.PermissionRepository.DeletePermissionByID(permissionID)
}

func (p PermissionDataService) UpdatePermission(permission *model.Permission) error {
	return p.PermissionRepository.UpdatePermission(permission)
}

func (p PermissionDataService) FindPermissionByID(permissionID int64) (*model.Permission, error) {
	return p.PermissionRepository.FindPermissionByID(permissionID)
}

func (p PermissionDataService) FindAllPermission() ([]model.Permission, error) {
	return p.PermissionRepository.FindAll()
}

func (p PermissionDataService) FindAllPermissionById(ids []int64) ([]*model.Permission, error) {
	return p.PermissionRepository.FindAllPermissionById(ids)
}
