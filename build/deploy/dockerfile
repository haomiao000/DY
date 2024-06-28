# 使用官方的Go语言基础镜像
FROM golang:1.17-alpine

# 设置工作目录
WORKDIR /app

# 将go.mod和go.sum复制到工作目录
COPY go.mod go.sum ./

# 下载并清理依赖
RUN go mod tidy

# 复制所有文件到工作目录
COPY . .

# 构建应用
RUN go build -o main .

# 暴露应用运行的端口
EXPOSE 8080

# 启动应用
CMD ["./main"]