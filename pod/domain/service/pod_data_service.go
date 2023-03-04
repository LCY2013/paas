package service

import (
	"context"
	"errors"
	log "github.com/LCY2013/paas/common/logger"
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/model"
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/repository"
	"github.com/LCY2013/thinking-in-go/paas/pod/proto/pod"
	"k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

// IPodDataService 这里是接口类型
type IPodDataService interface {
	AddPod(*model.Pod) (int64, error)
	DeletePod(int64) error
	UpdatePod(*model.Pod) error
	FindPodByID(int64) (*model.Pod, error)
	FindAllPod() ([]model.Pod, error)
	CreateToK8s(*pod.PodInfo) error
	DeleteFromK8s(*model.Pod) error
	UpdateToK8s(*pod.PodInfo) error
}

// NewPodDataService
// 创建
// 注意：返回值 IPodDataService 接口类型
func NewPodDataService(podRepository repository.IPodRepository, clientSet *kubernetes.Clientset) IPodDataService {
	return &PodDataService{PodRepository: podRepository, K8sClientSet: clientSet, deployment: &v1.Deployment{}}
}

type PodDataService struct {
	//注意：这里是 IPodRepository 类型
	PodRepository repository.IPodRepository
	K8sClientSet  *kubernetes.Clientset
	deployment    *v1.Deployment
}

// CreateToK8s 创建pod到k8s中
func (u *PodDataService) CreateToK8s(podInfo *pod.PodInfo) (err error) {
	u.SetDeployment(podInfo)
	if _, err = u.K8sClientSet.AppsV1().
		Deployments(podInfo.PodNamespace).
		Get(context.TODO(), podInfo.PodName, v12.GetOptions{}); err != nil {
		if _, err = u.K8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Create(context.TODO(), u.deployment, v12.CreateOptions{}); err != nil {
			log.Error(err)
			return err
		}
		log.Info("创建成功")
		return nil
	} else {
		//可以写自己的业务逻辑
		log.Error("Pod " + podInfo.PodName + "已经存在")
		return errors.New("Pod " + podInfo.PodName + " 已经存在")
	}

}

// UpdateToK8s 更新deployment，pod
func (u *PodDataService) UpdateToK8s(podInfo *pod.PodInfo) (err error) {
	u.SetDeployment(podInfo)
	if _, err = u.K8sClientSet.
		AppsV1().
		Deployments(podInfo.PodNamespace).
		Get(context.TODO(), podInfo.PodName, v12.GetOptions{}); err != nil {
		log.Error(err)
		return errors.New("Pod " + podInfo.PodName + " 不存在请先创建")
	} else {
		//如果存在
		if _, err = u.K8sClientSet.
			AppsV1().
			Deployments(podInfo.PodNamespace).
			Update(context.TODO(), u.deployment, v12.UpdateOptions{}); err != nil {
			log.Error(err)
			return err
		}
		log.Info(podInfo.PodName + " 更新成功")
		return nil
	}

}

// DeleteFromK8s 删除pod
func (u *PodDataService) DeleteFromK8s(pod *model.Pod) (err error) {
	if err = u.K8sClientSet.
		AppsV1().
		Deployments(pod.PodNamespace).
		Delete(context.TODO(), pod.PodName, v12.DeleteOptions{}); err != nil {
		log.Error(err)
		//写自己的业务逻辑
		return err
	} else {
		if err = u.DeletePod(pod.ID); err != nil {
			log.Error(err)
			return err
		}
		log.Info("删除Pod ID：" + strconv.FormatInt(pod.ID, 10) + " 成功！")
	}
	return
}

func (u *PodDataService) SetDeployment(podInfo *pod.PodInfo) {
	deployment := &v1.Deployment{}
	deployment.TypeMeta = v12.TypeMeta{
		Kind:       "deployment",
		APIVersion: "v1",
	}
	deployment.ObjectMeta = v12.ObjectMeta{
		Name:      podInfo.PodName,
		Namespace: podInfo.PodNamespace,
		Labels: map[string]string{
			"app-name": podInfo.PodName,
			"author":   "fufeng",
		},
	}
	deployment.Name = podInfo.PodName
	deployment.Spec = v1.DeploymentSpec{
		//副本个数
		Replicas: &podInfo.PodReplicas,
		Selector: &v12.LabelSelector{
			MatchLabels: map[string]string{
				"app-name": podInfo.PodName,
			},
			MatchExpressions: nil,
		},
		Template: v13.PodTemplateSpec{
			ObjectMeta: v12.ObjectMeta{
				Labels: map[string]string{
					"app-name": podInfo.PodName,
				},
			},
			Spec: v13.PodSpec{
				Containers: []v13.Container{
					{
						Name:            podInfo.PodName,
						Image:           podInfo.PodImage,
						Ports:           u.getContainerPort(podInfo),
						Env:             u.getEnv(podInfo),
						Resources:       u.getResources(podInfo),
						ImagePullPolicy: u.getImagePullPolicy(podInfo),
					},
				},
			},
		},
		Strategy:                v1.DeploymentStrategy{},
		MinReadySeconds:         0,
		RevisionHistoryLimit:    nil,
		Paused:                  false,
		ProgressDeadlineSeconds: nil,
	}
	u.deployment = deployment
}

func (u *PodDataService) getContainerPort(podInfo *pod.PodInfo) (containerPort []v13.ContainerPort) {
	for _, v := range podInfo.PodPort {
		containerPort = append(containerPort, v13.ContainerPort{
			Name:          "port-" + strconv.FormatInt(int64(v.ContainerPort), 10),
			ContainerPort: v.ContainerPort,
			Protocol:      u.getProtocol(v.Protocol),
		})
	}
	return
}

func (u *PodDataService) getProtocol(protocol string) v13.Protocol {
	switch protocol {
	case "TCP":
		return "TCP"
	case "UDP":
		return "UDP"
	case "SCTP":
		return "SCTP"
	default:
		return "TCP"
	}
}

func (u *PodDataService) getEnv(podInfo *pod.PodInfo) (envVar []v13.EnvVar) {
	for _, v := range podInfo.PodEnv {
		envVar = append(envVar, v13.EnvVar{
			Name:      v.EnvKey,
			Value:     v.EnvValue,
			ValueFrom: nil,
		})
	}
	return
}

func (u *PodDataService) getResources(podInfo *pod.PodInfo) (source v13.ResourceRequirements) {
	//最大能够使用多少资源
	source.Limits = v13.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(podInfo.PodCpuMax), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(podInfo.PodMemoryMax), 'f', 6, 64)),
	}
	//满足最少使用的资源量
	//@TODO 自己实现动态设置
	source.Requests = v13.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(podInfo.PodCpuMax), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(podInfo.PodMemoryMax), 'f', 6, 64)),
	}
	return
}

func (u *PodDataService) getImagePullPolicy(podInfo *pod.PodInfo) v13.PullPolicy {
	switch podInfo.PodPullPolicy {
	case "Always":
		return "Always"
	case "Never":
		return "Never"
	case "IfNotPresent":
		return "IfNotPresent"
	default:
		return "Always"
	}
}

// AddPod 插入
func (u *PodDataService) AddPod(pod *model.Pod) (int64, error) {
	return u.PodRepository.CreatePod(pod)
}

// DeletePod 删除
func (u *PodDataService) DeletePod(podID int64) error {
	return u.PodRepository.DeletePodByID(podID)
}

// UpdatePod 更新
func (u *PodDataService) UpdatePod(pod *model.Pod) error {
	return u.PodRepository.UpdatePod(pod)
}

// FindPodByID 查找
func (u *PodDataService) FindPodByID(podID int64) (*model.Pod, error) {
	return u.PodRepository.FindPodByID(podID)
}

// FindAllPod 查找
func (u *PodDataService) FindAllPod() ([]model.Pod, error) {
	return u.PodRepository.FindAll()
}
