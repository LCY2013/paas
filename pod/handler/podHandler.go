package handler

import (
	"context"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/service"
	pod "github.com/LCY2013/thinking-in-go/paas/pod/proto/pod"
)

type PodHandler struct {
	//注意这里的类型是 IPodDataService 接口类型
	PodDataService service.IPodDataService
}

// Call is a single request handler called via client.Call or the generated client code
func (e *PodHandler) AddPod(ctx context.Context, info *pod.PodInfo, rsp *pod.Response) error {
	log.Info("Received *pod.AddPod request")

	return nil
}

func (e *PodHandler) DeletePod(ctx context.Context, req *pod.PodId, rsp *pod.Response) error {
	log.Info("Received *pod.DeletePod request")

	return nil
}

func (e *PodHandler) UpdatePod(ctx context.Context, req *pod.PodInfo, rsp *pod.Response) error {
	log.Info("Received *pod.UpdatePod request")

	return nil
}

func (e *PodHandler) FindPodByID(ctx context.Context, req *pod.PodId, rsp *pod.PodInfo) error {
	log.Info("Received *pod.FindPodByID request")

	return nil
}

func (e *PodHandler) FindAllPod(ctx context.Context, req *pod.FindAll, rsp *pod.AllPod) error {
	log.Info("Received *pod.FindAllPod request")

	return nil
}
