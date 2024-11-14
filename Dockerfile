# 使用官方 Go 镜像作为构建环境
FROM golang:1.22-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的系统依赖
RUN apk add --no-cache git

# 首先只复制 go.mod
COPY go.mod ./

# 如果存在 go.sum，则复制它
COPY go.sum* ./

# 下载依赖
RUN go mod download || go mod tidy

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# 使用轻量级的 alpine 镜像作为运行环境
FROM alpine:latest

# 安装 ca-certificates，这对于 HTTPS 请求是必要的
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 创建日志目录
RUN mkdir -p /root/logs

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .
# 复制配置文件
COPY --from=builder /app/config/config.yaml ./config/config.yaml

# 设置环境变量
ENV CONFIG_PATH="config/config.yaml"

# 确保日志目录有正确的权限
RUN chmod 777 /root/logs

# 创建日志目录
VOLUME ["/root/logs"]

# 运行应用
CMD ["./main"] 