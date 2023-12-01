# 基础镜像
FROM debian:stable-slim AS prod
# 定义工作目录，用WORKDIR指定工作目录会在构建镜像的每一层中都存在，目录必须提前建好
WORKDIR /gcapi
RUN mkdir -p /gcapi/config/ && chmod 777 /gcapi
# 从上下文拷贝可执行文件到容器指定的目录，如果目录不存在则 COPY 命令会自动创建目录
COPY config/config.yaml /gcapi/config/config.yaml
COPY main /gcapi/gcapi
# 定义匿名卷，在启动容器时忘记挂载数据卷，会自动挂载到匿名卷：避免重要的数据因为容器重启而丢失，避免容器不断变大
VOLUME ["/gcapi"]
# 声明端口，定义该镜像的守护端口，以方便配置映射到物理机的端口
EXPOSE 8083
# CMD 在docker run 时运行，RUN是在docker build 时候执行
CMD ["/gcapi/gcapi", "--conf", "/gcapi/config/config.yaml"]