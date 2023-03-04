package repository

import (
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/svc/domain/model"
	"github.com/jinzhu/gorm"
)

// ISvcRepository 创建需要实现的接口
type ISvcRepository interface {
	// InitTable 初始化表
	InitTable() error
	// FindSvcByID 根据ID查处找数据
	FindSvcByID(int64) (*model.Svc, error)
	// CreateSvc 创建一条 svc 数据
	CreateSvc(*model.Svc) (int64, error)
	// DeleteSvcByID 根据ID删除一条 svc 数据
	DeleteSvcByID(int64) error
	// UpdateSvc 修改更新数据
	UpdateSvc(*model.Svc) error
	// FindAll 查找svc所有数据
	FindAll() ([]model.Svc, error)
}

// NewSvcRepository 创建svcRepository
func NewSvcRepository(db *gorm.DB) ISvcRepository {
	return &SvcRepository{mysqlDb: db}
}

type SvcRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *SvcRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Svc{}, &model.SvcPort{}).Error
}

// FindSvcByID 根据ID查找Svc信息
func (u *SvcRepository) FindSvcByID(svcID int64) (svc *model.Svc, err error) {
	svc = &model.Svc{}
	return svc, u.mysqlDb.First(svc, svcID).Error
}

// CreateSvc 创建Svc信息
func (u *SvcRepository) CreateSvc(svc *model.Svc) (int64, error) {
	return svc.ID, u.mysqlDb.Create(svc).Error
}

// DeleteSvcByID 根据ID删除Svc信息
func (u *SvcRepository) DeleteSvcByID(svcID int64) error {
	tx := u.mysqlDb.Begin()
	//遇到问题回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		log.Error(tx.Error)
		return tx.Error
	}
	//删除svc
	if err := u.mysqlDb.Where("id = ?", svcID).Delete(&model.Svc{}).Error; err != nil {
		tx.Rollback()
		log.Error(err)
		return err
	}
	//删除相关的port
	if err := u.mysqlDb.Where("svc_id = ?", svcID).Delete(&model.SvcPort{}).Error; err != nil {
		tx.Rollback()
		log.Error(err)
		return err
	}
	return tx.Commit().Error
}

// UpdateSvc 更新Svc信息
func (u *SvcRepository) UpdateSvc(svc *model.Svc) error {
	return u.mysqlDb.Model(svc).Update(svc).Error
}

// FindAll 获取结果集
func (u *SvcRepository) FindAll() (svcAll []model.Svc, err error) {
	return svcAll, u.mysqlDb.Find(&svcAll).Error
}
