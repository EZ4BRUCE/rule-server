FROM golang:alpine

LABEL MAINTAINER="EZ4BRUCE@lhy122786302@gmail.com"

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/build
WORKDIR /go/src/rule-server
# 将代码复制到容器中
COPY . .
RUN go mod tidy

# go generate 编译前自动执行代码
# go env 查看go的环境变量
# go build -o rule-server . 打包项目生成文件名为rule-server的二进制文件
RUN go generate && go env && go build -o rule-server .


FROM alpine:latest

RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

LABEL MAINTAINER="EZ4BRUCE@lhy122786302@gmail.com"
WORKDIR /go/src/rule-server

# 把/go/src/gin-vue-admin整个文件夹的文件到当前工作目录
COPY --from=0 /go/src/rule-server ./

EXPOSE 1016

ENTRYPOINT ./rule-server 