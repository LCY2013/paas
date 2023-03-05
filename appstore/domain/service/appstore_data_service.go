package service

import (
	"github.com/LCY2013/paas/appstore/domain/model"
	"github.com/LCY2013/paas/appstore/domain/repository"
	"k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

// IAppstoreDataService 这里是接口类型
type IAppstoreDataService interface {
	AddAppstore(*model.Appstore) (int64, error)
	DeleteAppstore(int64) error
	UpdateAppstore(*model.Appstore) error
	FindAppstoreByID(int64) (*model.Appstore, error)
	FindAllAppstore() ([]model.Appstore, error)

	//统计服务
	AddInstallNum(int64) error
	GetInstallNum(int64) int64
	AddViewNum(int64) error
	GetViewNum(int64) int64
}

// NewAppstoreDataService
// 创建
// 注意：返回值 IAppstoreDataService 接口类型
func NewAppstoreDataService(appstoreRepository repository.IAppstoreRepository, clientSet *kubernetes.Clientset) IAppstoreDataService {
	return &AppstoreDataService{AppstoreRepository: appstoreRepository, K8sClientSet: clientSet, deployment: &v1.Deployment{}}
}

type AppstoreDataService struct {
	//注意：这里是 IAppstoreRepository 类型
	AppstoreRepository repository.IAppstoreRepository
	K8sClientSet       *kubernetes.Clientset
	deployment         *v1.Deployment
}

// AddInstallNum 安装数量统计
func (u *AppstoreDataService) AddInstallNum(appID int64) error {
	return u.AppstoreRepository.AddInstallNumber(appID)
}

// GetInstallNum 查询安装数量
func (u *AppstoreDataService) GetInstallNum(appID int64) int64 {
	return u.AppstoreRepository.GetInstallNumber(appID)
}

// AddViewNum 添加浏览统计
func (u *AppstoreDataService) AddViewNum(appID int64) error {
	return u.AppstoreRepository.AddViewNumber(appID)
}

// GetViewNum 获取浏览量
func (u *AppstoreDataService) GetViewNum(appID int64) int64 {
	return u.AppstoreRepository.GetViewNumber(appID)
}

// AddAppstore 插入
func (u *AppstoreDataService) AddAppstore(appstore *model.Appstore) (int64, error) {
	return u.AppstoreRepository.CreateAppstore(appstore)
}

// DeleteAppstore 删除
func (u *AppstoreDataService) DeleteAppstore(appstoreID int64) error {
	return u.AppstoreRepository.DeleteAppstoreByID(appstoreID)
}

// UpdateAppstore 更新
func (u *AppstoreDataService) UpdateAppstore(appstore *model.Appstore) error {
	return u.AppstoreRepository.UpdateAppstore(appstore)
}

// FindAppstoreByID 查找
func (u *AppstoreDataService) FindAppstoreByID(appstoreID int64) (*model.Appstore, error) {
	return u.AppstoreRepository.FindAppstoreByID(appstoreID)
}

// FindAllAppstore 查找
func (u *AppstoreDataService) FindAllAppstore() ([]model.Appstore, error) {
	return u.AppstoreRepository.FindAll()
}
