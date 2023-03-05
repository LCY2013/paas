package service

import (
	"github.com/LCY2013/paas/user/domain/model"
	"github.com/LCY2013/paas/user/domain/repository"
	"k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

// IUserDataService 这里是接口类型
type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(*model.User) error
	FindUserByID(int64) (*model.User, error)
	FindAllUser() ([]model.User, error)

	// AddRole 分配角色
	AddRole(*model.User, []*model.Role) error
	// UpdateRole 更新用户角色
	UpdateRole(*model.User, []*model.Role) error
	// DeleteRole 删除用户角色
	DeleteRole(*model.User, []*model.Role) error

	// IsRight 判断用户是否有对应的权限
	IsRight(string, int64) bool
}

// NewUserDataService
// 创建
// 注意：返回值 IUserDataService 接口类型
func NewUserDataService(userRepository repository.IUserRepository, clientSet *kubernetes.Clientset) IUserDataService {
	return &UserDataService{UserRepository: userRepository, K8sClientSet: clientSet, deployment: &v1.Deployment{}}
}

type UserDataService struct {
	//注意：这里是 IUserRepository 类型
	UserRepository repository.IUserRepository
	K8sClientSet   *kubernetes.Clientset
	deployment     *v1.Deployment
}

func (u *UserDataService) AddRole(user *model.User, roles []*model.Role) error {
	return u.UserRepository.AddRole(user, roles)
}

func (u *UserDataService) UpdateRole(user *model.User, roles []*model.Role) error {
	return u.UserRepository.UpdateRole(user, roles)
}

func (u *UserDataService) DeleteRole(user *model.User, roles []*model.Role) error {
	return u.UserRepository.DeleteRole(user, roles)
}

func (u *UserDataService) IsRight(action string, userID int64) bool {
	return u.UserRepository.IsRight(action, userID)
}

// AddUser 插入
func (u *UserDataService) AddUser(user *model.User) (int64, error) {
	return u.UserRepository.CreateUser(user)
}

// DeleteUser 删除
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

// UpdateUser 更新
func (u *UserDataService) UpdateUser(user *model.User) error {
	return u.UserRepository.UpdateUser(user)
}

// FindUserByID 查找
func (u *UserDataService) FindUserByID(userID int64) (*model.User, error) {
	return u.UserRepository.FindUserByID(userID)
}

// FindAllUser 查找
func (u *UserDataService) FindAllUser() ([]model.User, error) {
	return u.UserRepository.FindAll()
}
