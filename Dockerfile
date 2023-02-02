#basic version
#run vs cmd vs entrypoint
#run: 在构建镜像的时候运行程序（在dockerfile里run）
#cmd：在docker利用镜像创建container的时候运行（例如默认启动命令）（可被exec执行命令覆盖）
#entrypoint：同cmd，但是不会被覆盖，必须先执行（用于比如数据库的参数创建）

# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
