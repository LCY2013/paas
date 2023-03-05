package handler

import (
	"context"
	"github.com/LCY2013/paas/user/domain/service"
	"github.com/LCY2013/paas/user/proto/permission"
)

type PermissionHandler struct {
	// PermissionDataService 注意这里的类型是 IPermissionDataService 接口类型
	PermissionDataService service.IPermissionDataService
}

func (p PermissionHandler) AddPermission(ctx context.Context, info *permission.PermissionInfo, response *permission.Response) error {
	//TODO implement me
	panic("implement me")
}

func (p PermissionHandler) DeletePermission(ctx context.Context, id *permission.PermissionId, response *permission.Response) error {
	//TODO implement me
	panic("implement me")
}

func (p PermissionHandler) UpdatePermission(ctx context.Context, info *permission.PermissionInfo, response *permission.Response) error {
	//TODO implement me
	panic("implement me")
}

func (p PermissionHandler) FindPermissionByID(ctx context.Context, id *permission.PermissionId, info *permission.PermissionInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p PermissionHandler) FindAllPermission(ctx context.Context, all *permission.FindAll, permission *permission.AllPermission) error {
	//TODO implement me
	panic("implement me")
}
