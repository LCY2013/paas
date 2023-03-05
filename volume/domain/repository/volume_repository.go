package repository

import (
	"github.com/LCY2013/paas/volume/domain/model"
	"github.com/jinzhu/gorm"
)

// IVolumeRepository 创建需要实现的接口
type IVolumeRepository interface {
	// InitTable 初始化表
	InitTable() error
	// FindVolumeByID 根据ID查处找数据
	FindVolumeByID(int64) (*model.Volume, error)
	// CreateVolume 创建一条 volume 数据
	CreateVolume(*model.Volume) (int64, error)
	// DeleteVolumeByID 根据ID删除一条 volume 数据
	DeleteVolumeByID(int64) error
	// UpdateVolume 修改更新数据
	UpdateVolume(*model.Volume) error
	// FindAll 查找volume所有数据
	FindAll() ([]model.Volume, error)
}

// NewVolumeRepository 创建volumeRepository
func NewVolumeRepository(db *gorm.DB) IVolumeRepository {
	return &VolumeRepository{mysqlDb: db}
}

type VolumeRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *VolumeRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Volume{}).Error
}

// FindVolumeByID 根据ID查找Volume信息
func (u *VolumeRepository) FindVolumeByID(volumeID int64) (volume *model.Volume, err error) {
	volume = &model.Volume{}
	return volume, u.mysqlDb.First(volume, volumeID).Error
}

// CreateVolume 创建Volume信息
func (u *VolumeRepository) CreateVolume(volume *model.Volume) (int64, error) {
	return volume.ID, u.mysqlDb.Create(volume).Error
}

// DeleteVolumeByID 根据ID删除Volume信息
func (u *VolumeRepository) DeleteVolumeByID(volumeID int64) error {
	return u.mysqlDb.Where("id = ?", volumeID).Delete(&model.Volume{}).Error
}

// UpdateVolume 更新Volume信息
func (u *VolumeRepository) UpdateVolume(volume *model.Volume) error {
	return u.mysqlDb.Model(volume).Update(volume).Error
}

// FindAll 获取结果集
func (u *VolumeRepository) FindAll() (volumeAll []model.Volume, err error) {
	return volumeAll, u.mysqlDb.Find(&volumeAll).Error
}
