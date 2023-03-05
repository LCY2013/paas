package handler

import (
	"context"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/user/domain/model"
	"github.com/LCY2013/paas/user/domain/service"
	"github.com/LCY2013/paas/user/proto/role"
)

type RoleHandler struct {
	// RoleDataService 注意这里的类型是 IRoleDataService 接口类型
	RoleDataService service.IRoleDataService

	PermissionDataService service.IPermissionDataService
}

func (r RoleHandler) getRolePermission(rolePermission *role.RolePermission) (role *model.Role, permission []*model.Permission, err error) {
	role, err = r.RoleDataService.FindRoleByID(rolePermission.RoleId)
	if err != nil {
		log.Error(err)
		return
	}

	permission, err = r.PermissionDataService.FindAllPermissionById(rolePermission.PermissionId)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (r RoleHandler) AddRole(ctx context.Context, info *role.RoleInfo, response *role.Response) error {
	//TODO implement me
	panic("implement me")
}

func (r RoleHandler) DeleteRole(ctx context.Context, id *role.RoleId, response *role.Response) error {
	//TODO implement me
	panic("implement me")
}

func (r RoleHandler) UpdateRole(ctx context.Context, info *role.RoleInfo, response *role.Response) error {
	//TODO implement me
	panic("implement me")
}

func (r RoleHandler) FindRoleByID(ctx context.Context, id *role.RoleId, info *role.RoleInfo) error {
	//TODO implement me
	panic("implement me")
}

func (r RoleHandler) FindAllRole(ctx context.Context, all *role.FindAll, role *role.AllRole) error {
	//TODO implement me
	panic("implement me")
}

func (r RoleHandler) AddPermission(ctx context.Context, rolePermission *role.RolePermission, response *role.Response) error {
	role, permission, err := r.getRolePermission(rolePermission)
	if err != nil {
		response.Msg = err.Error()
		return err
	}

	if err = r.RoleDataService.AddPermission(role, permission); err != nil {
		log.Error(err)
		response.Msg = err.Error()
		return err
	}

	return nil
}

func (r RoleHandler) UpdatePermission(ctx context.Context, rolePermission *role.RolePermission, response *role.Response) error {
	role, permission, err := r.getRolePermission(rolePermission)
	if err != nil {
		response.Msg = err.Error()
		return err
	}

	if err = r.RoleDataService.UpdatePermission(role, permission); err != nil {
		log.Error(err)
		response.Msg = err.Error()
		return err
	}

	return nil
}

func (r RoleHandler) DeletePermission(ctx context.Context, rolePermission *role.RolePermission, response *role.Response) error {
	role, permission, err := r.getRolePermission(rolePermission)
	if err != nil {
		response.Msg = err.Error()
		return err
	}

	if err = r.RoleDataService.DeletePermission(role, permission); err != nil {
		log.Error(err)
		response.Msg = err.Error()
		return err
	}

	return nil
}
