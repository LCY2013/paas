package service

import (
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/model"
	"github.com/LCY2013/thinking-in-go/paas/pod/domain/repository"
	"k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

// 这里是接口类型
type IPodDataService interface {
	AddPod(*model.Pod) (int64, error)
	DeletePod(int64) error
	UpdatePod(*model.Pod) error
	FindPodByID(int64) (*model.Pod, error)
	FindAllPod() ([]model.Pod, error)
}

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

// 插入
func (u *PodDataService) AddPod(pod *model.Pod) (int64, error) {
	return u.PodRepository.CreatePod(pod)
}

// 删除
func (u *PodDataService) DeletePod(podID int64) error {
	return u.PodRepository.DeletePodByID(podID)
}

// 更新
func (u *PodDataService) UpdatePod(pod *model.Pod) error {
	return u.PodRepository.UpdatePod(pod)
}

// 查找
func (u *PodDataService) FindPodByID(podID int64) (*model.Pod, error) {
	return u.PodRepository.FindPodByID(podID)
}

// 查找
func (u *PodDataService) FindAllPod() ([]model.Pod, error) {
	return u.PodRepository.FindAll()
}
