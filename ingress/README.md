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



```
