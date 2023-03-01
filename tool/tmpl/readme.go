package tmpl

var (
	Readme = `#  {{title .Alias}}

项目名称 {{title .Alias}}

# 1.快速创建代码仓库请使用下方命令
` + "```" + `
sudo docker run --rm -v $(pwd): $(pwd) -w  $(pwd) luochunyun/tool newService github.com/xxx/{{.Alias}}

注意：
1.sudo 如果是 Mac 系统提示输入的密码是你本机的密码。
2.这里github.com/xxx/{{.Alias}} 仓库名字需要和 go mod 名称一致
` + "```" + `


##  2.根据 proto 自动生成 go 基础代码
` + "```" + `
make proto
` + "```" + `

## 3.根据代码编译现有 Go 代码
` + "```" + `
make build
` + "```" + `
代码执行后会产生 {{.Alias}} 二进制文件

## 4.编译执行二进制文件
` + "```" + `
make docker
` + "```" + `
编译成功后会自动生成 {{.Alias}}:latest 镜像
可使用 docker images |grep {{.Alias}} 查看是否生成

## 5.该模版使用 go-micro v4 版本作为微服务开发框架
框架地址：https://github.com/asim/go-micro

## 6.开发流程
#### 开发前准备
* 使用tool镜像 (luochunyun/tool),项目工程目录生成工具。
* 使用 proto 生成工具镜像(luochunyun/protoc)，go-micro proto文件编码生成工具。
* micro-api-gateway(gharsallahmoez/micro)，根据go-micro v4 改造后的统一网关。

#### 1.开发服务端
* 1.1 使用 tool 完成项目工程目录生成
* 1.2 开发 domain - model
* 1.3 开发 domain - repository
* 1.4 开发 domain - service
* 1.5 开发 proto 文件，并且根据 proto 文件使用 luochunyun/protoc 工具生产代码
* 1.6 开发 对外暴露的服务
* 1.7 编写 main.go
* 1.8.1 打包 docker 镜像，编写 dockerfile（有操作k8s注意，拷贝或者挂载.kube/config 文件）
* 1.8.2 打包 docker 的时候注意注册中心，链路追踪地址要改成非内网地址，不然会造成访问失败 。
* 1.8.3 熔断器，监控系统的端口需要暴露和映射出去，不然收集不了数据。
* 1.8.4 服务队外的端口要固定micro.Address(":8081")。
* 1.8.5 自定义服务地址地址要修改正确，容器中修改成服务名称。
* 1.8.6 mysql 连接信息写上地址。
* 1.8.7 consul 中 mysql 连接地址需要修改。
* 1.8.8 交叉编译

` + "```" + `
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 替换成服务名称 *.go
` + "```" + `
* 1.8.9 编译

` + "```" + `
sudo docker build luochunyun/替换成服务名称
` + "```" + `
* 1.8.9 运行docker

` + "```" + `
sudo docker run  -p 8081:8081 -p 9092:9092 -p 9192:9192 /替换成绝对路径/.kube/config:/root/.kube/config luochunyun/替换成服务名称
` + "```" + `

#### 2.开发对外暴露的接口（启动API网关）
* 2.1 使用 tool 完成项目工程目录生成
* 2.2 编写 API proto 文件，并且根据 API proto 文件使用luochunyun/protoc 生成工具生成基础代码
* 2.3 编写对外暴露的 API 接口
* 2.4 编写 main.go
* 2.5 打包docker镜像


#### 3.启动网关
* 3.1使用 micro-api-gateway 建立网关

` + "```" + `
sudo docker run -d -p 8080:8080 gharsallahmoez/micro   --registry=consul --registry_address=替换成注册中心地址:8500 api --handler=api
` + "```" + `
注意：这里的地址要是能够访问的主机地址，不能是内网地址。

#### 4.开发前台应用


       
`

	ReadmeApi = `#  {{title .Alias}}

项目名称 {{title .Alias}}

# 1.快速创建代码仓库请使用下方命令
` + "```" + `
sudo docker run --rm -v $(pwd): $(pwd) -w  $(pwd) luochunyun/tool newService github.com/xxx/{{.Alias}}

注意：
1.sudo 如果是 Mac 系统提示输入的密码是你本机的密码。
2.这里github.com/xxx/{{.Alias}} 仓库名字需要和 go mod 名称一致
` + "```" + `


##  2.根据 proto 自动生成 go 基础代码
` + "```" + `
make proto
` + "```" + `

## 3.根据代码编译现有 Go 代码
` + "```" + `
make build
` + "```" + `
代码执行后会产生 {{.Alias}} 二进制文件

## 4.编译执行二进制文件
` + "```" + `
make docker
` + "```" + `
编译成功后会自动生成 {{.Alias}}:latest 镜像
可使用 docker images |grep {{.Alias}} 查看是否生成

## 5.该模版使用 go-micro v4 版本作为微服务开发框架
框架地址：https://github.com/asim/go-micro

## 6.开发流程
#### 开发前准备
* 使用tool镜像 (luochunyun/tool),项目工程目录生成工具。
* 使用 proto 生成工具镜像(luochunyun/protoc)，go-micro proto文件编码生成工具。
* micro-api-gateway(gharsallahmoez/micro)，根据go-micro v4 改造后的统一网关。

#### 1.开发服务端
* 1.1 使用 tool 完成项目工程目录生成
* 1.2 开发 domain - model
* 1.3 开发 domain - repository
* 1.4 开发 domain - service
* 1.5 开发 proto 文件，并且根据 proto 文件使用 luochunyun/protoc 工具生产代码
* 1.6 开发 对外暴露的服务
* 1.7 编写 main.go
* 1.8.1 打包 docker 镜像，编写 dockerfile（有操作k8s注意，拷贝或者挂载.kube/config 文件）
* 1.8.2 打包 docker 的时候注意注册中心，链路追踪地址要改成非内网地址，不然会造成访问失败 。
* 1.8.3 熔断器，监控系统的端口需要暴露和映射出去，不然收集不了数据。
* 1.8.4 服务队外的端口要固定micro.Address(":8081")。
* 1.8.5 自定义服务地址地址要修改正确，容器中修改成服务名称。
* 1.8.6 mysql 连接信息写上地址。
* 1.8.7 consul 中 mysql 连接地址需要修改。
* 1.8.8 交叉编译

` + "```" + `
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 替换成服务名称 *.go
` + "```" + `
* 1.8.9 编译

` + "```" + `
sudo docker build luochunyun/替换成服务名称
` + "```" + `
* 1.8.9 运行docker

` + "```" + `
sudo docker run  -p 8081:8081 -p 9092:9092 -p 9192:9192 /替换成绝对路径/.kube/config:/root/.kube/config luochunyun/替换成服务名称
` + "```" + `

#### 2.开发对外暴露的接口（启动API网关）
* 2.1 使用 tool 完成项目工程目录生成
* 2.2 编写 API proto 文件，并且根据 API proto 文件使用luochunyun/protoc 生成工具生成基础代码
* 2.3 编写对外暴露的 API 接口
* 2.4 编写 main.go
* 2.5 打包docker镜像


#### 3.启动网关
* 3.1使用 micro-api-gateway 建立网关

` + "```" + `
sudo docker run -d -p 8080:8080 gharsallahmoez/micro   --registry=consul --registry_address=替换成注册中心地址:8500 api --handler=api
` + "```" + `
注意：这里的地址要是能够访问的主机地址，不能是内网地址。

#### 4.开发前台应用


   
`
)
