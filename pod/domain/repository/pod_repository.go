package repository

import (
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/model"
	"github.com/jinzhu/gorm"
)

// 创建需要实现的接口
type IPodRepository interface {
	// InitTable 初始化表
	InitTable() error
	// FindPodByID 根据ID查处找数据
	FindPodByID(int64) (*model.Pod, error)
	// CreatePod 创建一条 pod 数据
	CreatePod(*model.Pod) (int64, error)
	// DeletePodByID 根据ID删除一条 pod 数据
	DeletePodByID(int64) error
	// UpdatePod 修改更新数据
	UpdatePod(*model.Pod) error
	// FindAll 查找pod所有数据
	FindAll() ([]model.Pod, error)
}

// NewPodRepository 创建podRepository
func NewPodRepository(db *gorm.DB) IPodRepository {
	return &PodRepository{mysqlDb: db}
}

type PodRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *PodRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Pod{}, &model.PodPort{}, &model.PodEnv{}).Error
}

// FindPodByID 根据ID查找Pod信息
func (u *PodRepository) FindPodByID(podID int64) (pod *model.Pod, err error) {
	pod = &model.Pod{}
	return pod, u.mysqlDb.
		Preload("PodPort").
		Preload("PodEnv").
		First(pod, podID).Error
}

// CreatePod 创建Pod信息
func (u *PodRepository) CreatePod(pod *model.Pod) (int64, error) {
	return pod.ID, u.mysqlDb.Create(pod).Error
}

// DeletePodByID 根据ID删除Pod信息
func (u *PodRepository) DeletePodByID(podID int64) error {
	tx := u.mysqlDb.Begin()
	//遇到问题回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	//彻底删除 POD 信息
	if err := u.mysqlDb.Where("id = ?", podID).Delete(&model.Pod{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//彻底删除 podenv 信息
	if err := u.mysqlDb.Where("pod_id = ?", podID).Delete(&model.PodEnv{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//彻底删除 podport 信息
	if err := u.mysqlDb.Where("pod_id = ?", podID).Delete(&model.PodPort{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdatePod 更新Pod信息
func (u *PodRepository) UpdatePod(pod *model.Pod) error {
	return u.mysqlDb.Model(pod).Update(pod).Error
}

// FindAll 获取结果集
func (u *PodRepository) FindAll() (podAll []model.Pod, err error) {
	return podAll, u.mysqlDb.Find(&podAll).Error
}
