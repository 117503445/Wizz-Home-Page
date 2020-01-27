FROM alpine

WORKDIR /web/gin

COPY ./out/linux/. .
COPY config.json .
EXPOSE 8080
CMD ./wizz-homepage-go
