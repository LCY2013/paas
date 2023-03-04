package handler

import (
	"context"
	"encoding/json"
	"errors"
	svc "github.com/LCY2013/paas/svc/proto/svc"
	"github.com/LCY2013/paas/svcapi/plugin/form"
	svcapi "github.com/LCY2013/paas/svcapi/proto/svcapi"
	log "go-micro.dev/v4/logger"
	"strconv"
)

type Svcapi struct {
	SvcService svc.SvcService
}

// FindSvcById svcapi.FindSvcById 通过API向外暴露为/svcapi/findSvcById，接收http请求
// 即：/svcapi/FindSvcById 请求会调用go.micro.api.svcapi 服务的svcapi.FindSvcById 方法
func (e *Svcapi) FindSvcById(ctx context.Context, req *svcapi.Request, rsp *svcapi.Response) error {
	log.Info("Received svcapi.FindSvcById request")
	if _, ok := req.Get["svc_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数异常")
	}
	//获取 svcId 参数
	svcIdString := req.Get["svc_id"].Values[0]
	svcId, err := strconv.ParseInt(svcIdString, 10, 64)
	if err != nil {
		log.Error(err)
		return err
	}
	//获取svc相关信息
	svcInfo, err := e.SvcService.FindSvcByID(ctx, &svc.SvcId{
		Id: svcId,
	})
	if err != nil {
		log.Error(err)
		return err
	}
	//json 返回svc信息

	rsp.StatusCode = 200
	b, _ := json.Marshal(svcInfo)
	rsp.Body = string(b)
	return nil
}

// AddSvc svcapi.AddSvc 通过API向外暴露为/svcapi/AddSvc，接收http请求
// 即：/svcapi/AddSvc 请求会调用go.micro.api.svcapi 服务的svcapi.AddSvc 方法
func (e *Svcapi) AddSvc(ctx context.Context, req *svcapi.Request, rsp *svcapi.Response) error {
	log.Info("Received svcapi.AddSvc request")
	//处理port
	addSvcInfo := &svc.SvcInfo{}
	svcType, ok := req.Post["svc_type"]
	if ok && len(svcType.Values) > 0 {
		svcPort := &svc.SvcPort{}
		switch svcType.Values[0] {
		case "ClusterIP":
			port, err := strconv.ParseInt(req.Post["svc_port"].Values[0], 10, 32)
			if err != nil {
				log.Error(err)
				return err
			}
			svcPort.SvcPort = int32(port)
			targetPort, err := strconv.ParseInt(req.Post["svc_target_port"].Values[0], 10, 32)
			if err != nil {
				log.Error(err)
				return err
			}
			svcPort.SvcTargetPort = int32(targetPort)
			svcPort.SvcPortProtocol = req.Post["svc_port_protocol"].Values[0]
			addSvcInfo.SvcPort = append(addSvcInfo.SvcPort, svcPort)
		default:
			return errors.New("类型不支持")
		}
	}
	//form 类型转换到结构体中
	form.FormToSvcStruct(req.Post, addSvcInfo)
	response, err := e.SvcService.AddSvc(ctx, addSvcInfo)
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	return nil
}

// DeleteSvcById svcapi.DeleteSvcById 通过API向外暴露为/svcapi/DeleteSvcById，接收http请求
// 即：/svcapi/DeleteSvcById 请求会调用go.micro.api.svcapi 服务的 svcapi.DeleteSvcById 方法
func (e *Svcapi) DeleteSvcById(ctx context.Context, req *svcapi.Request, rsp *svcapi.Response) error {
	log.Info("Received svcapi.DeleteSvcById request")
	if _, ok := req.Get["svc_id"]; !ok {
		return errors.New("参数异常")
	}
	//获取需要删除的ID
	svcIdString := req.Get["svc_id"].Values[0]
	svcId, err := strconv.ParseInt(svcIdString, 10, 64)
	if err != nil {
		log.Error(err)
		return err
	}
	//调用后端服务删除
	response, err := e.SvcService.DeleteSvc(ctx, &svc.SvcId{
		Id: svcId,
	})
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	return nil
}

// UpdateSvc svcapi.UpdateSvc 通过API向外暴露为/svcapi/UpdateSvc，接收http请求
// 即：/svcapi/UpdateSvc 请求会调用go.micro.api.svcapi 服务的svcapi.UpdateSvc 方法
func (e *Svcapi) UpdateSvc(ctx context.Context, req *svcapi.Request, rsp *svcapi.Response) error {
	log.Info("Received svcapi.UpdateSvc request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/svcApi/UpdateSvc'}")
	rsp.Body = string(b)
	return nil
}

// Call 默认的方法svcapi.Call 通过API向外暴露为/svcapi/call，接收http请求
// 即：/svcapi/call或/svcapi/ 请求会调用go.micro.api.svcapi 服务的svcapi.FindSvcById 方法
func (e *Svcapi) Call(ctx context.Context, req *svcapi.Request, rsp *svcapi.Response) error {
	log.Info("Received svcapi.Call request")
	allSvc, err := e.SvcService.FindAllSvc(ctx, &svc.FindAll{})
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(allSvc)
	rsp.Body = string(b)
	return nil
}
