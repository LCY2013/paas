package service

import (
	"github.com/LCY2013/paas/user/domain/model"
	"github.com/LCY2013/paas/user/domain/repository"
)

// IRoleDataService 这里是接口类型
type IRoleDataService interface {
	AddRole(role *model.Role) (int64, error)
	DeleteRole(roleID int64) error
	UpdateRole(role *model.Role) error
	FindRoleByID(roleID int64) (*model.Role, error)
	FindAllRole() ([]model.Role, error)

	// FindAllRoleByID 根据ID查找所有角色
	FindAllRoleByID(ids []int64) ([]*model.Role, error)

	// AddPermission 添加权限
	AddPermission(role *model.Role, permission []*model.Permission) error
	UpdatePermission(role *model.Role, permission []*model.Permission) error
	DeletePermission(role *model.Role, permission []*model.Permission) error
}

// NewRoleDataService
// 创建
// 注意：返回值 IRoleDataService 接口类型
func NewRoleDataService(roleRepository repository.IRoleRepository) IRoleDataService {
	return &RoleDataService{RoleRepository: roleRepository}
}

type RoleDataService struct {
	//注意：这里是 IRoleRepository 类型
	RoleRepository repository.IRoleRepository
}

func (r RoleDataService) AddRole(role *model.Role) (int64, error) {
	return r.RoleRepository.CreateRole(role)
}

func (r RoleDataService) DeleteRole(roleID int64) error {
	return r.RoleRepository.DeleteRoleByID(roleID)
}

func (r RoleDataService) UpdateRole(role *model.Role) error {
	return r.RoleRepository.UpdateRole(role)
}

func (r RoleDataService) FindRoleByID(roleID int64) (*model.Role, error) {
	return r.RoleRepository.FindRoleByID(roleID)
}

func (r RoleDataService) FindAllRole() ([]model.Role, error) {
	return r.RoleRepository.FindAll()
}

func (r RoleDataService) FindAllRoleByID(ids []int64) ([]*model.Role, error) {
	return r.RoleRepository.FindAllRoleById(ids)
}

func (r RoleDataService) AddPermission(role *model.Role, permission []*model.Permission) error {
	return r.RoleRepository.AddPermission(role, permission)
}

func (r RoleDataService) UpdatePermission(role *model.Role, permission []*model.Permission) error {
	return r.RoleRepository.UpdatePermission(role, permission)
}

func (r RoleDataService) DeletePermission(role *model.Role, permission []*model.Permission) error {
	return r.RoleRepository.DeletePermission(role, permission)
}
