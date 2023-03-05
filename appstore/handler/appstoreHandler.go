package handler

import (
	"context"
	"github.com/LCY2013/paas/appstore/domain/model"
	"github.com/LCY2013/paas/appstore/domain/service"
	appstore "github.com/LCY2013/paas/appstore/proto/appstore"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/common/util"
	"strconv"
)

type AppstoreHandler struct {
	// AppstoreDataService 注意这里的类型是 IAppstoreDataService 接口类型
	AppstoreDataService service.IAppstoreDataService
}

// 添加安装统计
func (e *AppstoreHandler) AddInstallNum(ctx context.Context, req *appstore.AppstoreId, rsp *appstore.Response) error {
	if err := e.AppstoreDataService.AddInstallNum(req.Id); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	rsp.Msg = "统计成功"
	return nil
}

// 获取安装数量
func (e *AppstoreHandler) GetInstallNum(ctx context.Context, req *appstore.AppstoreId, rsp *appstore.Number) error {
	rsp.Num = e.AppstoreDataService.GetInstallNum(req.Id)
	return nil
}

// 添加查询统计
func (e *AppstoreHandler) AddViewNum(ctx context.Context, req *appstore.AppstoreId, rsp *appstore.Response) error {
	if err := e.AppstoreDataService.AddViewNum(req.Id); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	rsp.Msg = "统计成功"
	return nil
}

// 获取查询数量
func (e *AppstoreHandler) GetViewNum(ctx context.Context, req *appstore.AppstoreId, rsp *appstore.Number) error {
	rsp.Num = e.AppstoreDataService.GetViewNum(req.Id)
	return nil
}

// AddAppstore Call is a single request handler called via client.Call or the generated client code
func (e *AppstoreHandler) AddAppstore(ctx context.Context, info *appstore.AppstoreInfo, rsp *appstore.Response) error {
	log.Info("Received *appstore.AddAppstore request")
	appStoreModel := &model.Appstore{}
	if err := util.SwapTo(info, appStoreModel); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}

	appStoreID, err := e.AppstoreDataService.AddAppstore(appStoreModel)
	if err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	rsp.Msg = "应用市场中新应用添加成功 ID 号为：" + strconv.FormatInt(appStoreID, 10)
	log.Info(rsp.Msg)
	return nil
}

func (e *AppstoreHandler) DeleteAppstore(ctx context.Context, req *appstore.AppstoreId, rsp *appstore.Response) error {
	log.Info("Received *appstore.DeleteAppstore request")
	return e.AppstoreDataService.DeleteAppstore(req.Id)
}

func (e *AppstoreHandler) UpdateAppstore(ctx context.Context, req *appstore.AppstoreInfo, rsp *appstore.Response) error {
	log.Info("Received *appstore.UpdateAppstore request")
	appStoreModel, err := e.AppstoreDataService.FindAppstoreByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if err := util.SwapTo(req, appStoreModel); err != nil {
		log.Error(err)
		return err
	}
	return e.AppstoreDataService.UpdateAppstore(appStoreModel)
}

func (e *AppstoreHandler) FindAppstoreByID(ctx context.Context, req *appstore.AppstoreId, rsp *appstore.AppstoreInfo) error {
	log.Info("Received *appstore.FindAppstoreByID request")
	appStoreModel, err := e.AppstoreDataService.FindAppstoreByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if err = util.SwapTo(appStoreModel, rsp); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *AppstoreHandler) FindAllAppstore(ctx context.Context, req *appstore.FindAll, rsp *appstore.AllAppstore) error {
	log.Info("Received *appstore.FindAllAppstore request")
	allAppStore, err := e.AppstoreDataService.FindAllAppstore()
	if err != nil {
		log.Error(err)
		return err
	}
	//整理数据格式
	for _, v := range allAppStore {
		appStoreInfo := &appstore.AppstoreInfo{}
		if err := util.SwapTo(v, appStoreInfo); err != nil {
			log.Error(err)
			return err
		}
		rsp.AppStoreInfo = append(rsp.AppStoreInfo, appStoreInfo)
	}
	return nil
}
