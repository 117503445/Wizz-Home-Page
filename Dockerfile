FROM golang as build

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release/Wizz-Home-Page

ADD . /go/release/Wizz-Home-Page

RUN go build -v -x -ldflags="-s -w" main.go

CMD ["./main"]