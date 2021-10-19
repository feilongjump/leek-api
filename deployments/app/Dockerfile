FROM golang:alpine

# 接收环境变量值
ARG TZ

# 为镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 使用国内的软件包镜像
RUN sed -i 's!https://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories

# 设置时区
RUN apk --no-cache add tzdata && \
    ln -snf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone

# 移动到工作目录：/code
WORKDIR /code

# 将代码复制到容器中
COPY ../../ .

# 下载依赖并且将代码编译成二进制可执行文件 leek_app
RUN go build -o leek_app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件 leek_app 从 /code 目录复制到这里
RUN cp /code/leek_app .

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
CMD ["/dist/leek_app"]