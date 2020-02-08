FROM golang as build

ENV GOPROXY=https://goproxy.cn GO111MODULE=on GIN_MODE=release PORT=80

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release/Wizz-Home-Page

ADD . /go/release/Wizz-Home-Page

RUN go build -v -x -ldflags="-s -w" main.go

CMD ["./main"]