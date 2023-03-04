## 版本使用 go-micro v4 

## consul 使用 ：1.15.0

不兼容 consul 1.10.以上版本，如果要兼容需要升级使用的 consul的依赖接口

## 说明
注释了其它插件的加载，如果需要修改需要自主安装

## 使用说明

```shell
go run main.go api \
  --registry=consul \
  --registry_address=192.168.0.57:8500 \
  --handler=api \
  --namespace=go.micro.api
  
go run main.go  api \
  --registry=consul \
  --registry_address=192.168.0.57:8500 \
  --enable_cors=true \
  --handler=api \
  --namespace=go.micro.api  
  
api --registry=consul --registry_address=192.168.0.57:8500 --enable_cors=true --handler=api --namespace=go.micro.api    
```
