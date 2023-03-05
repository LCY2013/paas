package handler

import (
	"context"
	"encoding/json"
	"errors"
	volume "github.com/LCY2013/paas/volume/proto/volume"
	"github.com/LCY2013/paas/volumeapi/plugin/form"
	volumeapi "github.com/LCY2013/paas/volumeapi/proto/volumeapi"
	log "go-micro.dev/v4/logger"
	"strconv"
)

type Volumeapi struct {
	VolumeService volume.VolumeService
}

// FindVolumeById volumeApi.FindVolumeById 通过API向外暴露为/volumeapi/findVolumeById，接收http请求
// 即：/volumeapi/FindVolumeById 请求会调用go.micro.api.volumeApi 服务的volumeApi.FindVolumeById 方法
func (e *Volumeapi) FindVolumeById(ctx context.Context, req *volumeapi.Request, rsp *volumeapi.Response) error {
	log.Info("Received volumeApi.FindVolumeById request")
	if _, ok := req.Get["volume_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数有异常")
	}
	//获取 volume_id
	volumeIdString := req.Get["volume_id"].Values[0]
	volumeId, err := strconv.ParseInt(volumeIdString, 10, 64)
	if err != nil {
		log.Error(err)
		return err
	}
	//获取 volume 信息
	volumeInfo, err := e.VolumeService.FindVolumeByID(ctx, &volume.VolumeId{
		Id: volumeId,
	})
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(volumeInfo)
	rsp.Body = string(b)
	return nil
}

// AddVolume volumeapi.AddVolume 通过API向外暴露为/volumeapi/AddVolume，接收http请求
// 即：/volumeapi/AddVolume 请求会调用go.micro.api.volumeapi 服务的volumeapi.AddVolume 方法
func (e *Volumeapi) AddVolume(ctx context.Context, req *volumeapi.Request, rsp *volumeapi.Response) error {
	log.Info("Received volumeApi.AddVolume request")
	addVolumeInfo := &volume.VolumeInfo{}
	form.FormToVolumeStruct(req.Post, addVolumeInfo)
	response, err := e.VolumeService.AddVolume(ctx, addVolumeInfo)
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	return nil
}

// DeleteVolumeById volumeApi.DeleteVolumeById 通过API向外暴露为/volumeapi/DeleteVolumeById，接收http请求
// 即：/volumeapi/DeleteVolumeById 请求会调用go.micro.api.volumeapi 服务的 volumeApi.DeleteVolumeById 方法
func (e *Volumeapi) DeleteVolumeById(ctx context.Context, req *volumeapi.Request, rsp *volumeapi.Response) error {
	log.Info("Received volumeApi.DeleteVolumeById request")
	if _, ok := req.Get["volume_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数异常")
	}
	//获取 volume_id
	volumeIdString := req.Get["volume_id"].Values[0]
	volumeId, err := strconv.ParseInt(volumeIdString, 10, 64)
	if err != nil {
		log.Error(err)
		return err
	}
	//调用 volume 删除服务
	response, err := e.VolumeService.DeleteVolume(ctx, &volume.VolumeId{
		Id: volumeId,
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

// UpdateVolume volumeapi.UpdateVolume 通过API向外暴露为/volumeapi/UpdateVolume，接收http请求
// 即：/volumeapi/UpdateVolume 请求会调用go.micro.api.volumeApi 服务的volumeapi.UpdateVolume 方法
func (e *Volumeapi) UpdateVolume(ctx context.Context, req *volumeapi.Request, rsp *volumeapi.Response) error {
	log.Info("Received volumeApi.UpdateVolume request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/volumeApi/UpdateVolume'}")
	rsp.Body = string(b)
	return nil
}

// Call 默认的方法volumeApi.Call 通过API向外暴露为/volumeapi/call，接收http请求
// 即：/volumeapi/call或/volumeApi/ 请求会调用go.micro.api.volumeapi 服务的volumeApi.FindVolumeById 方法
func (e *Volumeapi) Call(ctx context.Context, req *volumeapi.Request, rsp *volumeapi.Response) error {
	log.Info("Received volumeApi.Call request")
	allVolume, err := e.VolumeService.FindAllVolume(ctx, &volume.FindAll{})
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(allVolume)
	rsp.Body = string(b)
	return nil
}
