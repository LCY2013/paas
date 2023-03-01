# 在Docker中使用ProtoBuf

## Docker 根据.proto文件生成为服务文件

## proto-gen-go

### 一、构建镜像(dockerfile/protoc)

> docker build -f Dockerfile -t luochunyun/protoc:1.0.0 .

### 二、使用方法(proto目录)

```shell
protoc \
    -I proto \
    -I vendor/protoc-gen-validate \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    $(find proto -name '*.proto')
    
protoc \
    -I proto \
    -I vendor/protoc-gen-validate \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --micro_out=. \
    $(find ./ -name '*.proto')   
    
protoc --plugin=protoc-gen-go=$GOPATH/bin/protoc-gen-go --plugin=protoc-gen-micro=$GOPATH/bin/protoc-gen-micro --proto_path=. --micro_out=. --go_out=. greeter.proto     
```

> docker run --rm luochunyun/protoc:1.0.0 --version

> docker run --rm -v "$(pwd):/workspace" --workdir /workspace luochunyun/protoc:1.0.0 -I ./ --go_out=./ --micro_out=./ ./*.proto

> docker run --rm -v "$(pwd):/workspace" --workdir /workspace luochunyun/protoc:1.0.0 -I ./ --go_out=./ --micro_out=./ ./helloworld/helloworld.proto 

