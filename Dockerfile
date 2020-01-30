FROM golang:latest

#WORKDIR $GOPATH/src

#RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN export GO111MODULE=on

EXPOSE 8080

WORKDIR $GOPATH/src/Wizz-Home-Page

#RUN go get -v github.com/appleboy/gin-jwt/v2
#
#RUN go get -u -v github.com/gin-gonic/gin
#
#RUN go get -u github.com/go-sql-driver/mysql
#
#RUN go get -u github.com/jinzhu/gorm
#
#RUN go get -u github.com/mattn/go-sqlite3
#
#RUN go get -u github.com/spf13/viper



ADD . $GOPATH/src/Wizz-Home-Page

#RUN go get -v github.com/gin-gonic/gin
#RUN go get -v github.com/appleboy/gin-jwt/v2

RUN go mod download

CMD [ "go","run","main.go" ]