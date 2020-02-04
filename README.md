# Wizz-Home-Page

为之工作室 Wizz 的主页

## 加载配置

将 doc/config.json 复制到 项目根文件夹(和 main.go 同级),然后根据需要修改具体的配置项

account 为账号密码, 可以填写 {"admin": "admin","hello":"world"} 创建多个账号.因为偷懒所以都用了明文储存账号,请勿填写自己的常用密码!!!哪天有空再优化身份认证功能 :D

其中 图片 接口 需要填写 "qiniu" 有关字段.(因为七牛云有免费的 10 GB 空间) bucket 为空间名

事先向空间上传好 背景图片,然后把链接设置到 backgroundImageUrls 中

## 接口文档

本项目借助 swagger 进行接口文档托管

<http://ali.117503445.top:8080/swagger/index.html>

默认 host 是 ali.117503445.top:8080

### 修改方法

修改的话请修改 main.go 中的

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

#### 网络要求

默认启用国内镜像代理,如果服务器在国内就不用动

如果服务器在国外就编辑 Dockerfile 文件,删去

``` docker
 ENV GOPROXY https://goproxy.cn
```

这一行,取消代理

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

#### 构建镜像

- 下面 3 步 也可以通过运行 BuildScripts/DeployDocker.py 完成

- 在 Windows 上 ,下面 3 步 也可以通过运行 BuildScripts/DeployDocker.ps1 完成

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

## 运行时文件

运行时产生的文件都在 ./data 下

数据库文件为 Wizz-Home-Page.Database,采用 SQLite 3

日志在 ./data/logs 文件夹中,每次运行都会产生一个以运行时间命名的日志.

认证系统内部密码储存在 key.txt
