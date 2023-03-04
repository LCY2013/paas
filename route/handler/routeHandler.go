package handler

import (
	"context"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/common/util"
	"github.com/LCY2013/paas/route/domain/model"
	"github.com/LCY2013/paas/route/domain/service"
	route "github.com/LCY2013/paas/route/proto/route"
	"strconv"
)

type RouteHandler struct {
	// RouteDataService 注意这里的类型是 IRouteDataService 接口类型
	RouteDataService service.IRouteDataService
}

// AddRoute 添加路由
func (e *RouteHandler) AddRoute(ctx context.Context, info *route.RouteInfo, rsp *route.Response) error {
	log.Info("Received *route.AddRoute request")
	route := &model.Route{}
	if err := util.SwapTo(info, route); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	//创建route到k8s
	if err := e.RouteDataService.CreateRouteToK8s(info); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	} else {
		//写入数据库
		routeID, err := e.RouteDataService.AddRoute(route)
		if err != nil {
			log.Error(err)
			rsp.Msg = err.Error()
			return err
		}
		log.Info("Route 添加成功 ID 号为：" + strconv.FormatInt(routeID, 10))
		rsp.Msg = "Route 添加成功 ID 号为：" + strconv.FormatInt(routeID, 10)
	}
	return nil
}

// DeleteRoute 删除route
func (e *RouteHandler) DeleteRoute(ctx context.Context, req *route.RouteId, rsp *route.Response) error {
	log.Info("Received *route.DeleteRoute request")
	routeModel, err := e.RouteDataService.FindRouteByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	//从k8s中删除，并且删除数据库中数据
	if err = e.RouteDataService.DeleteRouteFromK8s(routeModel); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// UpdateRoute 更新route
func (e *RouteHandler) UpdateRoute(ctx context.Context, req *route.RouteInfo, rsp *route.Response) error {
	log.Info("Received *route.UpdateRoute request")
	if err := e.RouteDataService.UpdateRouteToK8s(req); err != nil {
		log.Error(err)
		return err
	}
	//查询数据库的信息
	routeModel, err := e.RouteDataService.FindRouteByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	//数据更新
	if err = util.SwapTo(req, routeModel); err != nil {
		log.Error(err)
		return err
	}
	return e.RouteDataService.UpdateRoute(routeModel)
}

// FindRouteByID 根据ID查询route信息
func (e *RouteHandler) FindRouteByID(ctx context.Context, req *route.RouteId, rsp *route.RouteInfo) error {
	log.Info("Received *route.FindRouteByID request")
	routeModel, err := e.RouteDataService.FindRouteByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	//数据转化
	if err = util.SwapTo(routeModel, rsp); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *RouteHandler) FindAllRoute(ctx context.Context, req *route.FindAll, rsp *route.AllRoute) error {
	log.Info("Received *route.FindAllRoute request")
	allRoute, err := e.RouteDataService.FindAllRoute()
	if err != nil {
		log.Error(err)
		return err
	}
	//整理下格式
	for _, v := range allRoute {
		//创建实例
		routeInfo := &route.RouteInfo{}
		//把查询出来的数据进行转化
		if err = util.SwapTo(v, routeInfo); err != nil {
			log.Error(err)
			return err
		}
		//数据合并
		rsp.RouteInfo = append(rsp.RouteInfo, routeInfo)
	}
	return nil
}
