FROM golang

WORKDIR $GOPATH/src/Wizz-Home-Page

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN export GO111MODULE=on

RUN go get github.com/appleboy/gin-jwt/v2

RUN go get github.com/gin-gonic/gin

RUN go get github.com/go-sql-driver/mysql

RUN go get github.com/jinzhu/gorm

RUN go get github.com/mattn/go-sqlite3

RUN go get github.com/spf13/viper

ADD . $GOPATH/src/Wizz-Home-Page

RUN go build .

EXPOSE 8080

CMD [ "go","run","main.go" ]