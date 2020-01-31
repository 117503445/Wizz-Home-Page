# Wizz-Home-Page

Wizz's Home Page

## 加载配置

将 doc/config.json 复制到 项目根文件夹(和 main.go 同级),然后根据需要修改具体的配置项

## 接口文档

本项目借助 swagger 进行接口文档托管

<http://ali.117503445.top:8080/swagger/index.html>

默认 host 是 ali.117503445.top:8080,如果要修改的话请修改 main.go 中的

```go
// @host ali.117503445.top:8080
```

然后执行 swag init,进行文档的更新

其中 swag 命令行工具如何获得我有空再写

## 部署方法

部署本项目可以通过 本机 go run 或者借助 Docker 进行部署

### 使用 go run (需要 Go 编译环境)

在项目根文件夹运行

```sh
go run main.go
```

### 使用 Docker 本地/远程 部署 (无需 Go 编译环境)

#### 主机要求

可能对内存有要求.本项目在 1.7G 的 Google Cloud 服务器上编译时出现内存不够用的错误,但是在 2G 的阿里云学生机上可以正常编译.

#### 准备本机 docker 环境

不管怎么样,需要在本地配置 docker 命令行工具

对于 Windows , 可以下载 <http://mirrors.aliyun.com/docker-toolbox/windows/docker-toolbox/>

#### 可选:连接远程服务器的 Docker server

先根据 <https://coreos.com/os/docs/latest/customizing-docker.html> 开启 Docker 远程访问

在本地环境变量中,添加

DOCKER_HOST

值为远程 docker server,比如

tcp://ali.117503445.top:2375

配置完成后,本地 docker 工具默认就是对远程 docker 进行操作

#### 编辑 Dockerfile 文件

如果服务器在国内就不用动

如果服务器在国外就删去

``` docker
 ENV GOPROXY https://goproxy.cn
```

这一行,取消代理

#### 构建镜像

在项目根文件夹下,键入

```docker
docker rm WizzHomePage -f
```

删除可能正在运行的 container

在项目根文件夹下,键入

```docker
docker build -t wizz .
```

构建名为 wizz 的 image

然后键入

```docker
docker run --name WizzHomePage -d -p 8080:8080 wizz
```

运行名为 WizzHomePage 的 CONTAINERS

在 Windows 上,上面 3 步 也可以通过运行 docker.ps1 完成
