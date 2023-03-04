package handler

import (
	"context"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/common/util"
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/model"
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/service"
	pod "github.com/LCY2013/thinking-in-go/paas/pod/proto/pod"
	"strconv"
)

type PodHandler struct {
	// PodDataService 注意这里的类型是 IPodDataService 接口类型
	PodDataService service.IPodDataService
}

// AddPod 添加创建POD
func (e *PodHandler) AddPod(ctx context.Context, info *pod.PodInfo, rsp *pod.Response) error {
	log.Info("Received *pod.AddPod request")
	podModel := &model.Pod{}
	err := util.SwapTo(info, podModel)
	if err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}

	if err = e.PodDataService.CreateToK8s(info); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	} else {
		//操作数据库写入数据
		podID, err := e.PodDataService.AddPod(podModel)
		if err != nil {
			log.Error(err)
			rsp.Msg = err.Error()
			return err
		}
		log.Info("Pod 添加成功数据库ID号为：" + strconv.FormatInt(podID, 10))
		rsp.Msg = "Pod 添加成功数据库ID号为：" + strconv.FormatInt(podID, 10)
	}
	return nil
}

// DeletePod 删除k8s中的pod 和数据库中的数据
func (e *PodHandler) DeletePod(ctx context.Context, req *pod.PodId, rsp *pod.Response) error {
	log.Info("Received *pod.DeletePod request")
	//先查找数据
	podModel, err := e.PodDataService.FindPodByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if err = e.PodDataService.DeleteFromK8s(podModel); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// UpdatePod 更新指定的pod
func (e *PodHandler) UpdatePod(ctx context.Context, req *pod.PodInfo, rsp *pod.Response) error {
	log.Info("Received *pod.UpdatePod request")
	//先更新k8s中的pod信息
	err := e.PodDataService.UpdateToK8s(req)
	if err != nil {
		log.Error(err)
		return err
	}
	//查询数据库中的pod
	podModel, err := e.PodDataService.FindPodByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	err = util.SwapTo(req, podModel)
	if err != nil {
		log.Error(err)
		return err
	}
	err = e.PodDataService.UpdatePod(podModel)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// FindPodByID 查询单个信息
func (e *PodHandler) FindPodByID(ctx context.Context, req *pod.PodId, rsp *pod.PodInfo) error {
	log.Info("Received *pod.FindPodByID request")
	//查询pod数据
	podModel, err := e.PodDataService.FindPodByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	err = util.SwapTo(podModel, rsp)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// FindAllPod 查询所有pod
func (e *PodHandler) FindAllPod(ctx context.Context, req *pod.FindAll, rsp *pod.AllPod) error {
	log.Info("Received *pod.FindAllPod request")
	allPod, err := e.PodDataService.FindAllPod()
	if err != nil {
		log.Error(err)
		return err
	}
	//整理格式
	for _, v := range allPod {
		podInfo := &pod.PodInfo{}
		err = util.SwapTo(v, podInfo)
		if err != nil {
			log.Error(err)
			return err
		}
		rsp.PodInfo = append(rsp.PodInfo, podInfo)
	}
	return nil
}
