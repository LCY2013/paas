# paas

## 业务代码目录

### 生成 service 代码

```shell
# 生成在当前目录的 pod service 目录
docker run --rm -v $(pwd):$(pwd) -w $(pwd) luochunyun/tool new github.com/LCY2013/paas/pod

# 生成完整目录的 pod service 目录
docker run --rm -v $(pwd):$(pwd) -w $(pwd) luochunyun/tool newService github.com/LCY2013/paas/pod
```

### 生成 api 代码

```shell
# 生成在当前目录的 pod api 目录
docker run --rm -v $(pwd):$(pwd) -w $(pwd) luochunyun/tool api github.com/LCY2013/paas/podapi

# 生成完整目录的 pod api 目录
docker run --rm -v $(pwd):$(pwd) -w $(pwd) luochunyun/tool apiCreate github.com/LCY2013/paas/pod
```

## proto 代码生成
```shell
cd 项目

make proto
```

## 项目构建
```shell
cd 项目

go mod tidy
```

## 镜像构建
```shell
make build-in-docker
```

## git 提交
```text
Feat：新功能（feature）
Fix：修补bug
Docs：文档（documentation）
Style： 格式（不影响代码运行的变动）
Refactor：重构（即不是新增功能，也不是修改bug的代码变动）
Test：增加测试
Chore：构建过程或辅助工具的变动
```
