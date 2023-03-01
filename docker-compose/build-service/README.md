## 通过docker-compose.yml构建镜像

```shell
docker-compose build base-service
```

## 通过docker-compose.yml启动镜像

```shell
docker-compose up -d
```

## 通过docker-compose.yml查看启动的镜像

```shell
# 通过 docker-compose 查看
docker-compose ps | grep base-service

# 通过 docker 查看
docker ps | grep base-service
```

## 通过docker-compose.yml停止镜像

```shell
docker-compose down
```
