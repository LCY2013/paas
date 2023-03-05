package handler

import (
	"context"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/common/util"
	"github.com/LCY2013/paas/volume/domain/model"
	"github.com/LCY2013/paas/volume/domain/service"
	volume "github.com/LCY2013/paas/volume/proto/volume"
	"strconv"
)

type VolumeHandler struct {
	// VolumeDataService 注意这里的类型是 IVolumeDataService 接口类型
	VolumeDataService service.IVolumeDataService
}

// AddVolume Call is a single request handler called via client.Call or the generated client code
func (e *VolumeHandler) AddVolume(ctx context.Context, info *volume.VolumeInfo, rsp *volume.Response) error {
	log.Info("Received *volume.AddVolume request")
	volume := &model.Volume{}
	if err := util.SwapTo(info, volume); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	//创建volume
	if err := e.VolumeDataService.CreateVolumeToK8s(info); err != nil {
		log.Error(err)
		rsp.Msg = err.Error()
		return err
	} else {
		//写入数据库
		volumeID, err := e.VolumeDataService.AddVolume(volume)
		if err != nil {
			log.Error(err)
			rsp.Msg = err.Error()
			return err
		}
		rsp.Msg = "volume 添加成功 ID 号为：" + strconv.FormatInt(volumeID, 10)
	}
	return nil
}

// DeleteVolume 删除
func (e *VolumeHandler) DeleteVolume(ctx context.Context, req *volume.VolumeId, rsp *volume.Response) error {
	log.Info("Received *volume.DeleteVolume request")
	volumModel, err := e.VolumeDataService.FindVolumeByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	//从k8s中删除，并且删除数据库
	if err := e.VolumeDataService.DeleteVolumeFromK8s(volumModel); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *VolumeHandler) UpdateVolume(ctx context.Context, req *volume.VolumeInfo, rsp *volume.Response) error {
	log.Info("Received *volume.UpdateVolume request")
	return nil
}

// FindVolumeByID 根据 ID 查找 volume
func (e *VolumeHandler) FindVolumeByID(ctx context.Context, req *volume.VolumeId, rsp *volume.VolumeInfo) error {
	log.Info("Received *volume.FindVolumeByID request")
	volumeModel, err := e.VolumeDataService.FindVolumeByID(req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	//数据转化
	if err = util.SwapTo(volumeModel, rsp); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *VolumeHandler) FindAllVolume(ctx context.Context, req *volume.FindAll, rsp *volume.AllVolume) error {
	log.Info("Received *volume.FindAllVolume request")
	allVolume, err := e.VolumeDataService.FindAllVolume()
	if err != nil {
		log.Error(err)
		return err
	}
	//整理格式
	for _, v := range allVolume {
		//创建实例
		volumeInfo := &volume.VolumeInfo{}
		//数据转化
		if err = util.SwapTo(v, volumeInfo); err != nil {
			log.Error(err)
			return err
		}
		//数据合并
		rsp.VolumeInfo = append(rsp.VolumeInfo, volumeInfo)
	}
	return nil
}
