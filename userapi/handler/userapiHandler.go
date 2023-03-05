package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/LCY2013/paas/user/proto/permission"
	"github.com/LCY2013/paas/user/proto/role"
	user "github.com/LCY2013/paas/user/proto/user"
	userapi "github.com/LCY2013/paas/userapi/proto/userapi"
	log "go-micro.dev/v4/logger"
	"strconv"
)

type Userapi struct {
	UserService       user.UserService
	RoleService       role.RoleService
	PermissionService permission.PermissionService
}

func (e *Userapi) getPost(req *userapi.Request, key string) (string, error) {
	if _, ok := req.Post[key]; !ok {
		return "", errors.New("参数不存在")
	}

	return req.Post[key].Values[0], nil
}

func (e *Userapi) getStringInt64(stringValue string) int64 {
	intValue, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		log.Error(err)
		return 0
	}
	return intValue
}

func (e *Userapi) AddRole(ctx context.Context, req *userapi.Request, resp *userapi.Response) error {
	userIdString, err := e.getPost(req, "user_id")
	if err != nil {
		return err
	}
	userId := e.getStringInt64(userIdString)

	if _, ok := req.Post["role_id"]; !ok {
		return errors.New("角色ID不存在")
	}

	var roleId []int64
	for _, v := range req.Post["role_id"].Values {
		roleId = append(roleId, e.getStringInt64(v))
	}

	rs, err := e.UserService.AddRole(ctx, &user.UserRole{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	resp.StatusCode = 200
	b, _ := json.Marshal(rs)
	resp.Body = string(b)
	return nil
}

func (e *Userapi) UpdateRole(ctx context.Context, req *userapi.Request, resp *userapi.Response) error {
	userIdString, err := e.getPost(req, "user_id")
	if err != nil {
		return err
	}
	userId := e.getStringInt64(userIdString)

	if _, ok := req.Post["role_id"]; !ok {
		return errors.New("角色ID不存在")
	}

	var roleId []int64
	for _, v := range req.Post["role_id"].Values {
		roleId = append(roleId, e.getStringInt64(v))
	}

	rs, err := e.UserService.UpdateRole(ctx, &user.UserRole{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	resp.StatusCode = 200
	b, _ := json.Marshal(rs)
	resp.Body = string(b)
	return nil
}

func (e *Userapi) DeleteRole(ctx context.Context, req *userapi.Request, resp *userapi.Response) error {
	userIdString, err := e.getPost(req, "user_id")
	if err != nil {
		return err
	}
	userId := e.getStringInt64(userIdString)

	if _, ok := req.Post["role_id"]; !ok {
		return errors.New("角色ID不存在")
	}

	var roleId []int64
	for _, v := range req.Post["role_id"].Values {
		roleId = append(roleId, e.getStringInt64(v))
	}

	rs, err := e.UserService.DeleteRole(ctx, &user.UserRole{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	resp.StatusCode = 200
	b, _ := json.Marshal(rs)
	resp.Body = string(b)
	return nil
}

func (e *Userapi) IsRight(ctx context.Context, req *userapi.Request, rsp *userapi.Response) error {
	if _, ok := req.Get["user_id"]; !ok {
		return errors.New("参数异常")
	}
	idString := req.Get["user_id"].Values[0]
	userId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Error(err)
		return err
	}

	if _, ok := req.Get["user_action"]; !ok {
		return errors.New("参数异常")
	}
	action := req.Get["user_action"].Values[0]
	if err != nil {
		log.Error(err)
		return err
	}
	right, err := e.UserService.IsRight(ctx, &user.UserRight{
		UserId: userId,
		Action: action,
	})

	if err != nil {
		log.Error(err)
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(right)
	rsp.Body = string(b)
	return nil
}

func (e *Userapi) AddPermission(ctx context.Context, req *userapi.Request, resp *userapi.Response) error {
	roleIdString, err := e.getPost(req, "role_id")
	if err != nil {
		return err
	}
	roleId := e.getStringInt64(roleIdString)

	if _, ok := req.Post["permission_id"]; !ok {
		return errors.New("角色ID不存在")
	}

	var permissionId []int64
	for _, v := range req.Post["permission_id"].Values {
		permissionId = append(permissionId, e.getStringInt64(v))
	}

	rs, err := e.RoleService.AddPermission(ctx, &role.RolePermission{
		RoleId:       roleId,
		PermissionId: permissionId,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	resp.StatusCode = 200
	b, _ := json.Marshal(rs)
	resp.Body = string(b)
	return nil
}

func (e *Userapi) UpdatePermission(ctx context.Context, req *userapi.Request, resp *userapi.Response) error {
	roleIdString, err := e.getPost(req, "role_id")
	if err != nil {
		return err
	}
	roleId := e.getStringInt64(roleIdString)

	if _, ok := req.Post["permission_id"]; !ok {
		return errors.New("角色ID不存在")
	}

	var permissionId []int64
	for _, v := range req.Post["permission_id"].Values {
		permissionId = append(permissionId, e.getStringInt64(v))
	}

	rs, err := e.RoleService.UpdatePermission(ctx, &role.RolePermission{
		RoleId:       roleId,
		PermissionId: permissionId,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	resp.StatusCode = 200
	b, _ := json.Marshal(rs)
	resp.Body = string(b)
	return nil
}

func (e *Userapi) DeletePermission(ctx context.Context, req *userapi.Request, resp *userapi.Response) error {
	roleIdString, err := e.getPost(req, "role_id")
	if err != nil {
		return err
	}
	roleId := e.getStringInt64(roleIdString)

	if _, ok := req.Post["permission_id"]; !ok {
		return errors.New("角色ID不存在")
	}

	var permissionId []int64
	for _, v := range req.Post["permission_id"].Values {
		permissionId = append(permissionId, e.getStringInt64(v))
	}

	rs, err := e.RoleService.DeletePermission(ctx, &role.RolePermission{
		RoleId:       roleId,
		PermissionId: permissionId,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	resp.StatusCode = 200
	b, _ := json.Marshal(rs)
	resp.Body = string(b)
	return nil
}

// userapi.FindUserById 通过API向外暴露为/userapi/findUserById，接收http请求
// 即：/userapi/FindUserById 请求会调用go.micro.api.userapi 服务的userapi.FindUserById 方法
func (e *Userapi) FindUserById(ctx context.Context, req *userapi.Request, rsp *userapi.Response) error {
	log.Info("Received userapi.FindUserById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/userapi/FindUserById'}")
	rsp.Body = string(b)
	return nil
}

// userapi.AddUser 通过API向外暴露为/userapi/AddUser，接收http请求
// 即：/userapi/AddUser 请求会调用go.micro.api.userapi 服务的userapi.AddUser 方法
func (e *Userapi) AddUser(ctx context.Context, req *userapi.Request, rsp *userapi.Response) error {
	log.Info("Received userapi.AddUser request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/userapi/AddUser'}")
	rsp.Body = string(b)
	return nil
}

// userapi.DeleteUserById 通过API向外暴露为/userapi/DeleteUserById，接收http请求
// 即：/userapi/DeleteUserById 请求会调用go.micro.api.userapi 服务的 userapi.DeleteUserById 方法
func (e *Userapi) DeleteUserById(ctx context.Context, req *userapi.Request, rsp *userapi.Response) error {
	log.Info("Received userapi.DeleteUserById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/userapi/DeleteUserById'}")
	rsp.Body = string(b)
	return nil
}

// userapi.UpdateUser 通过API向外暴露为/userapi/UpdateUser，接收http请求
// 即：/userapi/UpdateUser 请求会调用go.micro.api.userapi 服务的userapi.UpdateUser 方法
func (e *Userapi) UpdateUser(ctx context.Context, req *userapi.Request, rsp *userapi.Response) error {
	log.Info("Received userapi.UpdateUser request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/userapi/UpdateUser'}")
	rsp.Body = string(b)
	return nil
}

// 默认的方法userapi.Call 通过API向外暴露为/userapi/call，接收http请求
// 即：/userapi/call或/userapi/ 请求会调用go.micro.api.userapi 服务的userapi.FindUserById 方法
func (e *Userapi) Call(ctx context.Context, req *userapi.Request, rsp *userapi.Response) error {
	log.Info("Received userapi.Call request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问：Call'}")
	rsp.Body = string(b)
	return nil
}
