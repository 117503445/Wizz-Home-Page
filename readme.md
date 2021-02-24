# Wizz-Home-Page

## 本地开发环境部署

复制 ./config/config.toml.example 到 ./config/config.toml,再进行相应配置

借助 Goland 等 IDE,在 go run main.go 前先执行 preprocess.bat

## Docker生产环境部署

docker run --name wizz-home-page --restart=always -d -p 80:80 -v  ~/docker_volume/wizz-home-page/config:/root/config -v ~/docker_volume/wizz-home-page/tmp:/root/tmp 117503445/wizzhomepage

然后 复制配置文件到 ~/docker_volume/wizz-home-page/config/config.toml
