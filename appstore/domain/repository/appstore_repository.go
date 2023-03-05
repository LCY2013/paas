package repository

import (
	"github.com/LCY2013/paas/appstore/domain/model"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/jinzhu/gorm"
)

// IAppstoreRepository 创建需要实现的接口
type IAppstoreRepository interface {
	// InitTable 初始化表
	InitTable() error
	// FindAppstoreByID 根据ID查处找数据
	FindAppstoreByID(int64) (*model.Appstore, error)
	// CreateAppstore 创建一条 appstore 数据
	CreateAppstore(*model.Appstore) (int64, error)
	// DeleteAppstoreByID 根据ID删除一条 appstore 数据
	DeleteAppstoreByID(int64) error
	// UpdateAppstore 修改更新数据
	UpdateAppstore(*model.Appstore) error
	// FindAll 查找appstore所有数据
	FindAll() ([]model.Appstore, error)

	//添加安装数量
	AddInstallNumber(int64) error
	//获取安装数量
	GetInstallNumber(int64) int64
	//添加浏览量
	AddViewNumber(int64) error
	//获取浏览量
	GetViewNumber(int64) int64
}

// NewAppstoreRepository 创建appstoreRepository
func NewAppstoreRepository(db *gorm.DB) IAppstoreRepository {
	return &AppstoreRepository{mysqlDb: db}
}

type AppstoreRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *AppstoreRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Appstore{}, &model.AppComment{}, &model.AppVolume{}, &model.AppPod{}, &model.AppImage{}, &model.AppCategory{}, &model.AppIsv{}, &model.AppMiddle{}).Error
}

// 添加安装数量统计
func (u *AppstoreRepository) AddInstallNumber(appID int64) error {
	return u.mysqlDb.Model(&model.Appstore{}).Where("id = ?", appID).UpdateColumn("app_install", gorm.Expr("app_install + ?", 1)).Error
}

// 获取安装数量统计
func (u *AppstoreRepository) GetInstallNumber(appID int64) int64 {
	appStore, err := u.FindAppstoreByID(appID)
	if err != nil {
		log.Error(err)
		return 0
	}
	return appStore.AppInstall
}

// 添加浏览统计
func (u *AppstoreRepository) AddViewNumber(appID int64) error {
	return u.mysqlDb.Model(&model.Appstore{}).Where("id =  ?", appID).UpdateColumn("app_views", gorm.Expr("app_views + ?", 1)).Error
}

// 获取浏览数量
func (u *AppstoreRepository) GetViewNumber(appID int64) int64 {
	appStore, err := u.FindAppstoreByID(appID)
	if err != nil {
		log.Error(err)
		return 0
	}
	return appStore.AppViews
}

// FindAppstoreByID 根据ID查找Appstore信息
func (u *AppstoreRepository) FindAppstoreByID(appstoreID int64) (appstore *model.Appstore, err error) {
	appstore = &model.Appstore{}
	return appstore, u.mysqlDb.Preload("AppImage").Preload("AppPod").Preload("AppMiddle").Preload("AppVolume").Preload("AppComment").First(appstore, appstoreID).Error
}

// CreateAppstore 创建Appstore信息
func (u *AppstoreRepository) CreateAppstore(appstore *model.Appstore) (int64, error) {
	return appstore.ID, u.mysqlDb.Create(appstore).Error
}

// DeleteAppstoreByID 根据ID删除Appstore信息
func (u *AppstoreRepository) DeleteAppstoreByID(appstoreID int64) error {
	//开启事务
	tx := u.mysqlDb.Begin()
	//遇到问题回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	//遇到问题返回
	if tx.Error != nil {
		return tx.Error
	}

	//删除应用
	if err := u.mysqlDb.Where("id = ?", appstoreID).Delete(&model.Appstore{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除应用图片
	if err := u.mysqlDb.Where("app_id = ?", appstoreID).Delete(&model.AppImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除中间件
	if err := u.mysqlDb.Where("app_id = ?", appstoreID).Delete(&model.AppMiddle{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除对应的Pod组合
	if err := u.mysqlDb.Where("app_id = ?", appstoreID).Delete(&model.AppPod{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除存储
	if err := u.mysqlDb.Where("app_id = ?", appstoreID).Delete(&model.AppVolume{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除应用评论
	if err := u.mysqlDb.Where("app_id = ?", appstoreID).Delete(&model.AppComment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// UpdateAppstore 更新Appstore信息
func (u *AppstoreRepository) UpdateAppstore(appstore *model.Appstore) error {
	return u.mysqlDb.Model(appstore).Update(appstore).Error
}

// FindAll 获取结果集
func (u *AppstoreRepository) FindAll() (appstoreAll []model.Appstore, err error) {
	return appstoreAll, u.mysqlDb.Find(&appstoreAll).Error
}
