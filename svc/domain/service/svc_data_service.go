package service

import (
	"context"
	"errors"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/paas/svc/domain/model"
	"github.com/LCY2013/paas/svc/domain/repository"
	"github.com/LCY2013/paas/svc/proto/svc"
	v1Dep "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

// ISvcDataService 这里是接口类型
type ISvcDataService interface {
	AddSvc(*model.Svc) (int64, error)
	DeleteSvc(int64) error
	UpdateSvc(*model.Svc) error
	FindSvcByID(int64) (*model.Svc, error)
	FindAllSvc() ([]model.Svc, error)
	CreateSvcToK8s(*svc.SvcInfo) error
	UpdateSvcToK8s(*svc.SvcInfo) error
	DeleteFromK8s(*model.Svc) error
}

// NewSvcDataService
// 创建
// 注意：返回值 ISvcDataService 接口类型
func NewSvcDataService(svcRepository repository.ISvcRepository, clientSet *kubernetes.Clientset) ISvcDataService {
	return &SvcDataService{SvcRepository: svcRepository, K8sClientSet: clientSet, deployment: &v1Dep.Deployment{}}
}

type SvcDataService struct {
	//注意：这里是 ISvcRepository 类型
	SvcRepository repository.ISvcRepository
	K8sClientSet  *kubernetes.Clientset
	deployment    *v1Dep.Deployment
}

// 创建服务到K8s中
func (u *SvcDataService) CreateSvcToK8s(svcInfo *svc.SvcInfo) (err error) {
	service := u.setService(svcInfo)
	//查找是否纯在指定的服务
	if _, err = u.K8sClientSet.
		CoreV1().
		Services(svcInfo.SvcNamespace).
		Get(context.TODO(), svcInfo.SvcName, v12.GetOptions{}); err != nil {
		//查找不到,就创建
		if _, err = u.K8sClientSet.
			CoreV1().
			Services(svcInfo.SvcNamespace).
			Create(context.TODO(), service, v12.CreateOptions{}); err != nil {
			log.Error(err)
			return err
		}
		return nil
	} else {
		log.Error("Service " + svcInfo.SvcName + "已经存在")
		return errors.New("Service " + svcInfo.SvcName + "已经存在")
	}

}

// 根据svcnfo 设置Iservice 信息
func (u *SvcDataService) setService(svcInfo *svc.SvcInfo) *v1.Service {
	service := &v1.Service{}
	if svcInfo.SvcType == "" {
		svcInfo.SvcType = "ClusterIP"
	}
	//设置服务类型
	service.TypeMeta = v12.TypeMeta{
		Kind:       "v1",
		APIVersion: "Service",
	}
	//设置服务基础信息
	service.ObjectMeta = v12.ObjectMeta{
		Name:      svcInfo.SvcName,
		Namespace: svcInfo.SvcNamespace,
		Labels: map[string]string{
			"app-name": svcInfo.SvcPodName,
			"author":   "fufeng",
		},
		Annotations: map[string]string{
			"k8s/generated-by-fufeng": "create by fufeng",
		},
	}
	//设置服务的spec信息，采用ClusterIP模式
	service.Spec = v1.ServiceSpec{
		Ports: u.getSvcPort(svcInfo),
		Selector: map[string]string{
			"app-name": svcInfo.SvcPodName,
		},
		//Type: "ClusterIP",
		Type: v1.ServiceType(svcInfo.SvcType),
	}
	return service
}

func (u *SvcDataService) getSvcPort(svcInfo *svc.SvcInfo) (servicePort []v1.ServicePort) {
	for _, v := range svcInfo.SvcPort {
		servicePort = append(servicePort, v1.ServicePort{
			Name:       "port-" + strconv.FormatInt(int64(v.SvcPort), 10),
			Protocol:   v1.Protocol(v.SvcPortProtocol),
			Port:       v.SvcPort,
			TargetPort: intstr.FromInt(int(v.SvcTargetPort)),
		})
	}
	return
}

func (u *SvcDataService) UpdateSvcToK8s(svcInfo *svc.SvcInfo) (err error) {
	service := u.setService(svcInfo)
	//查找是否纯在指定的服务
	if _, err = u.K8sClientSet.
		CoreV1().
		Services(svcInfo.SvcNamespace).
		Get(context.TODO(), svcInfo.SvcName, v12.GetOptions{}); err != nil {
		//查找不到
		log.Error(err)
		return errors.New("Service" + svcInfo.SvcName + "不存在请先创建")
	} else {
		if _, err = u.K8sClientSet.
			CoreV1().
			Services(svcInfo.SvcNamespace).
			Update(context.TODO(), service, v12.UpdateOptions{}); err != nil {
			log.Error(err)
			return err
		}
		log.Info("Service " + svcInfo.SvcName + "更新成功")
		return nil
	}
}

func (u *SvcDataService) DeleteFromK8s(svc *model.Svc) (err error) {
	if err = u.K8sClientSet.
		CoreV1().
		Services(svc.SvcNamespace).
		Delete(context.TODO(), svc.SvcName, v12.DeleteOptions{}); err != nil {
		log.Error(err)
		return err
	} else {
		if err := u.DeleteSvc(svc.ID); err != nil {
			log.Error(err)
			return err
		}
		log.Info("删除Service ID：" + strconv.FormatInt(svc.ID, 10) + "成功！")
	}
	return
}

// AddSvc 插入
func (u *SvcDataService) AddSvc(svc *model.Svc) (int64, error) {
	return u.SvcRepository.CreateSvc(svc)
}

// DeleteSvc 删除
func (u *SvcDataService) DeleteSvc(svcID int64) error {
	return u.SvcRepository.DeleteSvcByID(svcID)
}

// UpdateSvc 更新
func (u *SvcDataService) UpdateSvc(svc *model.Svc) error {
	return u.SvcRepository.UpdateSvc(svc)
}

// FindSvcByID 查找
func (u *SvcDataService) FindSvcByID(svcID int64) (*model.Svc, error) {
	return u.SvcRepository.FindSvcByID(svcID)
}

// FindAll 查找
func (u *SvcDataService) FindAllSvc() ([]model.Svc, error) {
	return u.SvcRepository.FindAll()
}
