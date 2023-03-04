package repository

import (
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/route/domain/model"
	"github.com/jinzhu/gorm"
)

// IRouteRepository 创建需要实现的接口
type IRouteRepository interface {
	// InitTable 初始化表
	InitTable() error
	// FindRouteByID 根据ID查处找数据
	FindRouteByID(int64) (*model.Route, error)
	// CreateRoute 创建一条 route 数据
	CreateRoute(*model.Route) (int64, error)
	// DeleteRouteByID 根据ID删除一条 route 数据
	DeleteRouteByID(int64) error
	// UpdateRoute 修改更新数据
	UpdateRoute(*model.Route) error
	// FindAll 查找route所有数据
	FindAll() ([]model.Route, error)
}

// NewRouteRepository 创建routeRepository
func NewRouteRepository(db *gorm.DB) IRouteRepository {
	return &RouteRepository{mysqlDb: db}
}

type RouteRepository struct {
	mysqlDb *gorm.DB
}

// InitTable 初始化表
func (u *RouteRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Route{}, &model.RoutePath{}).Error
}

// FindRouteByID 根据ID查找Route信息
func (u *RouteRepository) FindRouteByID(routeID int64) (route *model.Route, err error) {
	route = &model.Route{}
	return route, u.mysqlDb.Preload("RoutePath").First(route, routeID).Error
}

// CreateRoute 创建Route信息
func (u *RouteRepository) CreateRoute(route *model.Route) (int64, error) {
	return route.ID, u.mysqlDb.Create(route).Error
}

// DeleteRouteByID 根据ID删除Route信息
func (u *RouteRepository) DeleteRouteByID(routeID int64) error {
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
	//开始删除
	if err := u.mysqlDb.Where("id = ?", routeID).Delete(&model.Route{}).Error; err != nil {
		tx.Rollback()
		log.Error(err)
		return err
	}
	//删除关联表
	if err := u.mysqlDb.Where("route_id = ?", routeID).Delete(&model.RoutePath{}).Error; err != nil {
		tx.Rollback()
		log.Error(err)
		return err
	}
	return tx.Commit().Error
}

// UpdateRoute 更新Route信息
func (u *RouteRepository) UpdateRoute(route *model.Route) error {
	return u.mysqlDb.Model(route).Update(route).Error
}

// FindAll 获取结果集
func (u *RouteRepository) FindAll() (routeAll []model.Route, err error) {
	return routeAll, u.mysqlDb.Preload("RoutePath").Find(&routeAll).Error
}
