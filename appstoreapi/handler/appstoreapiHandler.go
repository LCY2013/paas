package handler

import (
	"context"
	"encoding/json"
	"errors"
	appstore "github.com/LCY2013/paas/appstore/proto/appstore"
	"github.com/LCY2013/paas/appstoreapi/plugin/form"
	appstoreapi "github.com/LCY2013/paas/appstoreapi/proto/appstoreapi"
	log "go-micro.dev/v4/logger"
	"strconv"
)

type Appstoreapi struct {
	AppstoreService appstore.AppstoreService
}

// GetId 获取 url 中的应用ID
func (e *Appstoreapi) GetId(req *appstoreapi.Request) (int64, error) {
	if _, ok := req.Get["app_id"]; !ok {
		return 0, errors.New("参数异常")
	}
	//获取到ID后进行转化
	IdString := req.Get["app_id"].Values[0]
	Id, err := strconv.ParseInt(IdString, 10, 64)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return Id, nil
}

// appstoreapi.FindAppstoreById 通过API向外暴露为/appstoreapi/findAppstoreById，接收http请求
// 即：/appstoreapi/FindAppstoreById 请求会调用go.micro.api.appstoreapi 服务的appstoreapi.FindAppstoreById 方法
func (e *Appstoreapi) FindAppstoreById(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	log.Info("Received appstoreapi.FindAppstoreById request")
	Id, err := e.GetId(req)
	if err != nil {
		log.Error(err)
		return err
	}
	//获取应用市场中应用的相关信息
	info, err := e.AppstoreService.FindAppstoreByID(ctx, &appstore.AppstoreId{
		Id: Id,
	})
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(info)
	rsp.Body = string(b)
	return nil
}

// appstoreapi.AddAppstore 通过API向外暴露为/appstoreapi/AddAppstore，接收http请求
// 即：/appstoreapi/AddAppstore 请求会调用go.micro.api.appstoreapi 服务的appstoreapi.AddAppstore 方法
func (e *Appstoreapi) AddAppstore(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	log.Info("Received appstoreapi.AddAppstore request")
	addAppStore := &appstore.AppstoreInfo{}
	//进行简单form数据映射
	form.FormToAppStoreStruct(req.Post, addAppStore)
	//设置图片
	e.SetImage(req, addAppStore)
	//设置POD
	e.SetPod(req, addAppStore)
	//设置中间件
	e.SetMiddle(req, addAppStore)
	//设置存储
	e.SetVolume(req, addAppStore)

	//调用后端服务进行更新
	response, err := e.AppstoreService.AddAppstore(ctx, addAppStore)
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	return nil
}

// 设置图片
func (e *Appstoreapi) SetImage(req *appstoreapi.Request, appStoreInfo *appstore.AppstoreInfo) {
	dataSlice, ok := req.Post["app_image"]
	if ok {
		imageSlice := []*appstore.AppImage{}
		for _, v := range dataSlice.Values {
			image := &appstore.AppImage{
				AppImageSrc: v,
			}
			imageSlice = append(imageSlice, image)
		}
		appStoreInfo.AppImage = imageSlice
	}
}

// 设置POD模板
func (e *Appstoreapi) SetPod(req *appstoreapi.Request, appStoreInfo *appstore.AppstoreInfo) {
	dataSlice, ok := req.Post["app_pod"]
	if ok {
		var podSlice []*appstore.AppPod
		for _, v := range dataSlice.Values {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Error(err)
				continue
			}
			pod := &appstore.AppPod{
				AppPodId: id,
			}
			podSlice = append(podSlice, pod)
		}
		appStoreInfo.AppPod = podSlice
	}
}

// SetMiddle 设置中间件模板
func (e *Appstoreapi) SetMiddle(req *appstoreapi.Request, appStoreInfo *appstore.AppstoreInfo) {
	dataSlice, ok := req.Post["app_middle"]
	if ok {
		var middleSlice []*appstore.AppMiddle
		for _, v := range dataSlice.Values {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Error(err)
				continue
			}
			middle := &appstore.AppMiddle{
				AppMiddleId: id,
			}
			middleSlice = append(middleSlice, middle)
		}
		appStoreInfo.AppMiddle = middleSlice
	}

}

// 设置存储
func (e *Appstoreapi) SetVolume(req *appstoreapi.Request, appStoreInfo *appstore.AppstoreInfo) {
	dataSlice, ok := req.Post["app_volume"]
	if ok {
		var volumeSlice []*appstore.AppVolume
		for _, v := range dataSlice.Values {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Error(err)
				continue
			}
			volume := &appstore.AppVolume{
				AppVolumeId: id,
			}
			volumeSlice = append(volumeSlice, volume)
		}
		appStoreInfo.AppVolume = volumeSlice
	}
}

// appstoreapi.DeleteAppstoreById 通过API向外暴露为/appstoreapi/DeleteAppstoreById，接收http请求
// 即：/appstoreapi/DeleteAppstoreById 请求会调用go.micro.api.appstoreapi 服务的 appstoreapi.DeleteAppstoreById 方法
func (e *Appstoreapi) DeleteAppstoreById(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	log.Info("Received appstoreapi.DeleteAppstoreById request")
	Id, err := e.GetId(req)
	if err != nil {
		log.Error(err)
		return err
	}
	response, err := e.AppstoreService.DeleteAppstore(ctx, &appstore.AppstoreId{
		Id: Id,
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

// appstoreapi.UpdateAppstore 通过API向外暴露为/appstoreapi/UpdateAppstore，接收http请求
// 即：/appstoreapi/UpdateAppstore 请求会调用go.micro.api.appstoreapi 服务的appstoreapi.UpdateAppstore 方法
func (e *Appstoreapi) UpdateAppstore(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	log.Info("Received appstoreapi.UpdateAppstore request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/appstoreapi/UpdateAppstore'}")
	rsp.Body = string(b)
	return nil
}

// 默认的方法appstoreapi.Call 通过API向外暴露为/appstoreapi/call，接收http请求
// 即：/appstoreapi/call或/appstoreapi/ 请求会调用go.micro.api.appstoreapi 服务的appstoreapi.FindAppstoreById 方法
func (e *Appstoreapi) Call(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	log.Info("Received appstoreapi.Call request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问：Call'}")
	rsp.Body = string(b)
	return nil
}

// AddInstallNum 安装统计接口
func (e *Appstoreapi) AddInstallNum(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	Id, err := e.GetId(req)
	if err != nil {
		log.Error(err)
		return err
	}
	response, err := e.AppstoreService.AddInstallNum(ctx, &appstore.AppstoreId{
		Id: Id,
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

// GetInstallNum 获取安装数量
func (e *Appstoreapi) GetInstallNum(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	Id, err := e.GetId(req)
	if err != nil {
		log.Error(err)
		return err
	}
	response, err := e.AppstoreService.GetInstallNum(ctx, &appstore.AppstoreId{
		Id: Id,
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

// AddViewNum 安装统计接口
func (e *Appstoreapi) AddViewNum(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	Id, err := e.GetId(req)
	if err != nil {
		log.Error(err)
		return err
	}
	response, err := e.AppstoreService.AddViewNum(ctx, &appstore.AppstoreId{
		Id: Id,
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

// GetViewNum 获取安装数量
func (e *Appstoreapi) GetViewNum(ctx context.Context, req *appstoreapi.Request, rsp *appstoreapi.Response) error {
	Id, err := e.GetId(req)
	if err != nil {
		log.Error(err)
		return err
	}
	response, err := e.AppstoreService.GetViewNum(ctx, &appstore.AppstoreId{
		Id: Id,
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
