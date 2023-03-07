# go-micro-template

基于 Go 的微服务架构后端模版。主服务使用 Gin 向外提供 API，各服务间通过 gRPC 进行通信。预置 Mysql 与 Redis 用作数据存储。

## CLI 工具

- [**Docker**](https://hub.docker.com/)
- [**golang-migrate**](https://github.com/golang-migrate/migrate)（数据库迁移）
- [**sqlc**](https://docs.sqlc.dev/en/stable/index.html)（sql->go 接口）
- [**Protocol Buffers**](https://protobuf.dev)（gRPC 数据定义）
- [**Evans**](https://github.com/ktr0731/evans)（gRPC 调试工具）

## 主要依赖

- [**gRPC**](https://grpc.io/)
- [**Protocol Buffers**](https://protobuf.dev)（gRPC 数据定义）
- [**Asynq**](https://github.com/hibiken/asynq)（任务队列异步处理框架）
- [**golang-migrate**](https://github.com/golang-migrate/migrate)（数据库迁移）
- [**Testify**](https://github.com/stretchr/testify)（测试框架）
- [**Viper**](https://github.com/spf13/viper)（配置项管理）
- [**Paseto**](https://github.com/o1egl/paseto)（用户鉴权）
- [**Zerolog**](https://github.com/rs/zerolog)（日志输出）

## 项目修改

- 修改 go.mod 中 module 名（全局替换）
- 向 .gitignore 文件中添加 app.env 与 compose.yaml

## 开发场景

#### 基本环境

- 安装 CLI 工具
- 安装 Go 依赖 `go mod tidy`

#### 数据库

- Mysql

  - 执行`migrate_init`生成 schema
  - 使用 [**dbdiagram**](https://dbdiagram.io/home) 工具设计数据库，将 sql 语句复制到上一步的 schema 中
  - 执行`make mysql DB_PASSWORD=? DB_NAME=?`，启动 mysql 容器

- Redis

  - 执行`make redis`，启动 mysql 容器

#### sqlc

- 在 internal/db/query/ 下创建 表名.sql 文件，根据官网编写 sql 语句
- 执行`make sqlc`生成.go 文件

#### protoc

- 修改 internal/proto 目录下的文件名以及文件内容
- 执行 `make proto`

#### 编译运行

- 执行`make server`，编译运行

## 部署场景

- 修改 compose.yaml 与根目录下的 makefile 相关信息
- 将项目文件直接上传至服务器后，在服务器端打包镜像`make build_images`

#### Swarm 方式部署（支持多节点集群）

> 参考 [**docker swarm**](https://docs.docker.com/engine/reference/commandline/swarm/)、[**docker service**](https://docs.docker.com/engine/reference/commandline/service/)、[**docker stack**](https://docs.docker.com/engine/reference/commandline/stack/)

- 执行`docker swarm init`，创建集群
- 执行`docker stack deploy -c compose.yaml ???`（??? 为项目名），部署 stack
- 执行`docker service ls`，查看 service 列表
- 执行`docker service logs SERVICE`，查看某个 service 的日志

- 执行`docker stack rm ???`，结束 stack
- 执行`docker swarm leave`，离开集群

#### Compose 方式部署（仅支持单节点）

> 参考 [**docker compose**](https://docs.docker.com/engine/reference/commandline/compose/)

- 执行`docker compose up -d`，后台运行
- 执行`docker compose logs`，查看日志

- 执行`docker compose down`，结束运行并删除容器
