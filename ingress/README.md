#### 在master节点执行命令
```shell
kubectl apply -f deploy.yml
```

### 查询SVC
```shell
kubectl get svc

ingress-nginx-controller             ClusterIP   10.96.0.117   <none>        80/TCP,443/TCP   9s
ingress-nginx-controller-admission   ClusterIP   10.96.0.195   <none>        443/TCP          9s
kubernetes                           ClusterIP   10.96.0.1     <none>        443/TCP          3d20h
nginx                                ClusterIP   10.96.2.45    <none>        80/TCP           145m

------------------------------------------------------------------------------------------------

kubectl get IngressClass

NAME    CONTROLLER             PARAMETERS   AGE
nginx   k8s.io/ingress-nginx   <none>       5m59s

------------------------------------------------------------------------------------------------

kubectl get daemonset

NAME                       DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
ingress-nginx-controller   2         2         2       2            2           kubernetes.io/os=linux   73s

kubectl get ing
NAME          CLASS   HOSTS         ADDRESS       PORTS   AGE
nginx-route   nginx   nginx-route.com   10.96.3.196   80      72s


kubectl get po -o wide
NAME                                   READY   STATUS      RESTARTS   AGE     IP              NODE         NOMINATED NODE   READINESS GATES
nginx-68b884cdc8-jvv4g                 1/1     Running     0          4h47m   100.93.142.8    k8snode202   <none>           <none>
nginx-68b884cdc8-lw6kn                 1/1     Running     0          4h47m   100.66.222.11   k8snode201   <none>           <none>
```

### 不需要修改hosts

```shell
curl --resolve nginx-route.com:80:192.168.0.201 "http://nginx-route.com/"

curl --resolve nginx-route.com:80:192.168.0.202 "http://nginx-route.com/"
```
