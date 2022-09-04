FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/xiamei.guo/blog-api
COPY .. $GOPATH/src/xiamei.guo/blog-api
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./blog-api"]