FROM golang:1.16 as build
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/release
ADD . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag i
RUN GOOS=linux go build -ldflags="-s -w" -o app
FROM alpine as prod
EXPOSE 80
COPY --from=build /go/release/app /app
WORKDIR /
ENTRYPOINT /app