FROM alpine

WORKDIR /web/gin

COPY ./out/linux/. .

CMD ./wizz-homepage-go
