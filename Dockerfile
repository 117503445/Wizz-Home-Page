FROM golang

WORKDIR $GOPATH/src/Wizz-Home-Page

ADD . $GOPATH/src/Wizz-Home-Page

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN export GO111MODULE=on

RUN go get github.com/appleboy/gin-jwt/v2

RUN go build .

EXPOSE 8080

CMD [ "go","run","main.go" ]