package repository

import (
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/model"
	"github.com/jinzhu/gorm"
)

// 创建需要实现的接口
type IPodRepository interface {
	//初始化表
	InitTable() error
	//根据ID查处找数据
	FindPodByID(int64) (*model.Pod, error)
	//创建一条 pod 数据
	CreatePod(*model.Pod) (int64, error)
	//根据ID删除一条 pod 数据
	DeletePodByID(int64) error
	//修改更新数据
	UpdatePod(*model.Pod) error
	//查找pod所有数据
	FindAll() ([]model.Pod, error)
}

// 创建podRepository
func NewPodRepository(db *gorm.DB) IPodRepository {
	return &PodRepository{mysqlDb: db}
}

type PodRepository struct {
	mysqlDb *gorm.DB
}

// 初始化表
func (u *PodRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Pod{}).Error
}

// 根据ID查找Pod信息
func (u *PodRepository) FindPodByID(podID int64) (pod *model.Pod, err error) {
	pod = &model.Pod{}
	return pod, u.mysqlDb.First(pod, podID).Error
}

// 创建Pod信息
func (u *PodRepository) CreatePod(pod *model.Pod) (int64, error) {
	return pod.ID, u.mysqlDb.Create(pod).Error
}

// 根据ID删除Pod信息
func (u *PodRepository) DeletePodByID(podID int64) error {
	return u.mysqlDb.Where("id = ?", podID).Delete(&model.Pod{}).Error
}

// 更新Pod信息
func (u *PodRepository) UpdatePod(pod *model.Pod) error {
	return u.mysqlDb.Model(pod).Update(pod).Error
}

// 获取结果集
func (u *PodRepository) FindAll() (podAll []model.Pod, err error) {
	return podAll, u.mysqlDb.Find(&podAll).Error
}
