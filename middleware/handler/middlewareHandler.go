package handler

import (
	"context"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/common/util"
	"github.com/LCY2013/paas/middleware/domain/model"
	"github.com/LCY2013/paas/middleware/domain/service"
	middleware "github.com/LCY2013/paas/middleware/proto/middleware"
	"strconv"
)

type MiddlewareHandler struct {
	// MiddlewareDataService 注意这里的类型是 IMiddlewareDataService 接口类型
	MiddlewareDataService service.IMiddlewareDataService
	// 添加中间件类型服务
	MiddleTypeDataService service.IMiddleTypeDataService
}

// AddMiddleware Call is a single request handler called via client.Call or the generated client code
func (e *MiddlewareHandler) AddMiddleware(ctx context.Context, info *middleware.MiddlewareInfo, rsp *middleware.Response) error {
	log.Info("Received *middleware.AddMiddleware request")

	middleModel := &model.Middleware{}
	if err := util.SwapTo(info, middleModel); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	//调用其它的服务处理数据
	//根据ID产销需要的镜像地址
	imageAddress, err := e.MiddleTypeDataService.FindImageVersionByID(info.MiddleVersionId)
	if err != nil {
		log.Error(err)
		return err
	}
	//赋值
	info.MiddleDockerImageVersion = imageAddress
	//在k8s 中创建资源
	if err = e.MiddlewareDataService.CreateToK8s(info); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	} else {
		//插入数据库
		middleID, err := e.MiddlewareDataService.AddMiddleware(middleModel)
		if err != nil {
			log.Error(err)
			rsp.Msg = err.Error()
			return err
		}
		rsp.Msg = "中间件添加成功 ID 号为：" + strconv.FormatInt(middleID, 10)
		log.Info(rsp.Msg)
	}

	return nil
}

func (e *MiddlewareHandler) DeleteMiddleware(ctx context.Context, req *middleware.MiddlewareId, rsp *middleware.Response) error {
	log.Info("Received *middleware.DeleteMiddleware request")
	middleModel, err := e.MiddlewareDataService.FindMiddlewareByID(req.Id)
	if err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	//删除k8s中资源
	if err = e.MiddlewareDataService.DeleteFromK8s(middleModel); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

func (e *MiddlewareHandler) UpdateMiddleware(ctx context.Context, req *middleware.MiddlewareInfo, rsp *middleware.Response) error {
	log.Info("Received *middleware.UpdateMiddleware request")
	if err := e.MiddlewareDataService.UpdateToK8s(req); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	//查询中间件相关的信息
	middleModle, err := e.MiddlewareDataService.FindMiddlewareByID(req.Id)
	if err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	//更新model数据
	if err = util.SwapTo(req, middleModle); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	//更新model
	if err = e.MiddlewareDataService.UpdateMiddleware(middleModle); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

func (e *MiddlewareHandler) FindMiddlewareByID(ctx context.Context, req *middleware.MiddlewareId, rsp *middleware.MiddlewareInfo) error {
	log.Info("Received *middleware.FindMiddlewareByID request")
	//查询中间件
	middlewareModel, err := e.MiddlewareDataService.FindMiddlewareByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if err = util.SwapTo(middlewareModel, rsp); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *MiddlewareHandler) FindAllMiddleware(ctx context.Context, req *middleware.FindAll, rsp *middleware.AllMiddleware) error {
	log.Info("Received *middleware.FindAllMiddleware request")

	allMiddleware, err := e.MiddlewareDataService.FindAllMiddleware()
	if err != nil {
		log.Error(err)
		return err
	}
	//整理格式
	for _, v := range allMiddleware {
		middleInfo := &middleware.MiddlewareInfo{}
		if err = util.SwapTo(v, middleInfo); err != nil {
			log.Error(err)
			return err
		}
		rsp.MiddlewareInfo = append(rsp.MiddlewareInfo, middleInfo)
	}

	return nil
}

// 查找所有的中间件
func (e *MiddlewareHandler) FindAllMiddlewareByTypeID(ctx context.Context, req *middleware.FindAllByTypeId, rsp *middleware.AllMiddleware) error {
	log.Info("Received *middleware.FindAllMiddleware request")
	allMiddleware, err := e.MiddlewareDataService.FindAllMiddlewareByTypeID(req.TypeId)
	if err != nil {
		log.Error(err)
		return err
	}
	//整理格式
	for _, v := range allMiddleware {
		middleInfo := &middleware.MiddlewareInfo{}
		if err := util.SwapTo(v, middleInfo); err != nil {
			log.Error(err)
			return err
		}
		rsp.MiddlewareInfo = append(rsp.MiddlewareInfo, middleInfo)
	}
	return nil
}

// 根据ID查找中间件类型信息
func (e *MiddlewareHandler) FindMiddleTypeByID(ctx context.Context, req *middleware.MiddleTypeId, rsp *middleware.MiddleTypeInfo) error {
	typeModel, err := e.MiddleTypeDataService.FindMiddleTypeByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if err := util.SwapTo(typeModel, rsp); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// 添加中间件
func (e *MiddlewareHandler) AddMiddleType(ctx context.Context, info *middleware.MiddleTypeInfo, rsp *middleware.Response) error {
	typeModel := &model.MiddleType{}
	if err := util.SwapTo(info, typeModel); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	id, err := e.MiddleTypeDataService.AddMiddleType(typeModel)
	if err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	rsp.Msg = "中间件类型 ID 号为： " + strconv.FormatInt(id, 10)
	log.Info(rsp.Msg)
	return nil
}

// 删除中间件类型
func (e *MiddlewareHandler) DeleteMiddleTypeByID(ctx context.Context, req *middleware.MiddleTypeId, rsp *middleware.Response) error {
	if err := e.MiddleTypeDataService.DeleteMiddleType(req.Id); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

// 更新中间件类型
func (e *MiddlewareHandler) UpdateMiddleType(ctx context.Context, req *middleware.MiddleTypeInfo, rsp *middleware.Response) error {
	typeModel, err := e.MiddleTypeDataService.FindMiddleTypeByID(req.Id)
	if err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	if err := util.SwapTo(req, typeModel); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	if err := e.MiddleTypeDataService.UpdateMiddleType(typeModel); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

// 查找所有的类型
func (e *MiddlewareHandler) FindAllMiddleType(ctx context.Context, req *middleware.FindAll, rsp *middleware.AllMiddleType) error {
	//查询所有中间件
	allMiddleType, err := e.MiddleTypeDataService.FindAllMiddleType()
	if err != nil {
		log.Error(err)
		return err
	}
	//整理格式
	for _, v := range allMiddleType {
		middleInfo := &middleware.MiddleTypeInfo{}
		if err := util.SwapTo(v, middleInfo); err != nil {
			log.Error(err)
			return err
		}
		rsp.MiddleTypeInfo = append(rsp.MiddleTypeInfo, middleInfo)
	}
	return nil
}
