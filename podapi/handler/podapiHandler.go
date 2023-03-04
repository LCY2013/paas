package handler

import (
	"context"
	"encoding/json"
	"errors"
	pod "github.com/LCY2013/paas/pod/proto/pod"
	"github.com/LCY2013/paas/podapi/plugin/form"
	podapi "github.com/LCY2013/paas/podapi/proto/podapi"
	log "go-micro.dev/v4/logger"
	"strconv"
)

type Podapi struct {
	PodService pod.PodService
}

// FindPodById
// podapi.FindPodById 通过API向外暴露为/podapi/findPodById，接收http请求
// 即：/podapi/FindPodById 请求会调用go.micro.api.podapi 服务的podapi.FindPodById 方法
func (e *Podapi) FindPodById(ctx context.Context, req *podapi.Request, rsp *podapi.Response) error {
	log.Info("Received podapi.FindPodById request")
	if _, ok := req.Get["pod_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数异常")
	}
	//获取podid 参数
	podIdString := req.Get["pod_id"].Values[0]
	podId, err := strconv.ParseInt(podIdString, 10, 64)
	if err != nil {
		return err
	}
	//获取pod相关信息
	podInfo, err := e.PodService.FindPodByID(ctx, &pod.PodId{
		Id: podId,
	})
	if err != nil {
		return err
	}
	//json 返回pod信息
	rsp.StatusCode = 200
	b, _ := json.Marshal(podInfo)
	rsp.Body = string(b)
	return nil
}

// AddPod
// podapi.AddPod 通过API向外暴露为/podapi/AddPod，接收http请求
// 即：/podapi/AddPod 请求会调用go.micro.api.podapi 服务的podapi.AddPod 方法
func (e *Podapi) AddPod(ctx context.Context, req *podapi.Request, rsp *podapi.Response) error {
	log.Info("Received podapi.AddPod request")
	addPodInfo := &pod.PodInfo{}
	//处理 port
	dataSlice, ok := req.Post["pod_port"]
	if ok {
		//特殊处理
		var podSlice []*pod.PodPort
		for _, v := range dataSlice.Values {
			i, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				log.Error(err)
			}
			port := &pod.PodPort{
				ContainerPort: int32(i),
				Protocol:      "TCP",
			}
			podSlice = append(podSlice, port)
		}
		addPodInfo.PodPort = podSlice
	}
	//form类型转化到结构体中
	form.FromToPodStruct(req.Post, addPodInfo)

	response, err := e.PodService.AddPod(ctx, addPodInfo)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(response)
	rsp.Body = string(b)
	return nil
}

// DeletePodById
// podapi.DeletePodById 通过API向外暴露为/podapi/DeletePodById，接收http请求
// 即：/podapi/DeletePodById 请求会调用go.micro.api.podapi 服务的 podapi.DeletePodById 方法
func (e *Podapi) DeletePodById(ctx context.Context, req *podapi.Request, rsp *podapi.Response) error {
	log.Info("Received podapi.DeletePodById request")
	if _, ok := req.Get["pod_id"]; !ok {
		return errors.New("参数异常")
	}
	//获取要删除的ID
	podIdString := req.Get["pod_id"].Values[0]
	podId, err := strconv.ParseInt(podIdString, 10, 64)
	if err != nil {
		log.Error(err)
		return err
	}
	//删除指定服务
	response, err := e.PodService.DeletePod(ctx, &pod.PodId{
		Id: podId,
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

// UpdatePod
// podapi.UpdatePod 通过API向外暴露为/podapi/UpdatePod，接收http请求
// 即：/podapi/UpdatePod 请求会调用go.micro.api.podapi 服务的podapi.UpdatePod 方法
func (e *Podapi) UpdatePod(ctx context.Context, req *podapi.Request, rsp *podapi.Response) error {
	log.Info("Received podapi.UpdatePod request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/podapi/UpdatePod'}")
	rsp.Body = string(b)
	return nil
}

// Call
// 默认的方法podapi.Call 通过API向外暴露为/podapi/call，接收http请求
// 即：/podapi/call或/podapi/ 请求会调用go.micro.api.podapi 服务的podapi.FindPodById 方法
func (e *Podapi) Call(ctx context.Context, req *podapi.Request, rsp *podapi.Response) error {
	log.Info("Received podapi.Call request")
	allPod, err := e.PodService.FindAllPod(ctx, &pod.FindAll{})
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(allPod)
	rsp.Body = string(b)
	return nil
}
