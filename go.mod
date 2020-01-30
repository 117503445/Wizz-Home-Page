module Wizz-Home-Page

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/appleboy/gin-jwt/v2 v2.6.3
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.12
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	github.com/spf13/viper v1.6.2
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
)

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.52.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200128174031-69ecbb4d6d5d
	golang.org/x/net => github.com/golang/net v0.0.0-20200114155413-6afb5195e5aa
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200124204421-9fbb57f87de9
	golang.org/x/text => github.com/golang/text v0.3.2
	google.golang.org/appengine => github.com/golang/appengine v1.6.5
)
