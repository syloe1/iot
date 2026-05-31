# IoT Platform

基于 go-zero 的 IoT 学习项目，采用 monorepo 结构，当前已经实现用户、产品、设备、开放平台、设备消息下发、设备状态消费和后台状态推送的基础链路。

## 技术栈

- Go
- go-zero
- gRPC / zrpc
- MySQL
- Redis
- Etcd
- EMQX / MQTT
- SSE

## 项目结构

```text
app/
  user/
    api/          用户 HTTP API
    rpc/          用户 RPC，登录/注册
  product/
    api/          产品 HTTP API
    rpc/          产品 RPC，CRUD/列表
  device/
    api/          设备 HTTP API，CRUD/列表
    rpc/          设备 RPC，CRUD/列表/下发消息
    consumer/     MQTT 消费服务，处理设备状态/心跳
  openapi/
    api/          开放平台 API，签名验证 + 下发设备消息
  push/
    api/          后台实时推送 API，SSE 推送设备状态事件
common/
  mqttx/          MQTT client helper
  redisx/         Redis client helper
  util/           JWT helper
```

## 已实现功能

- 用户注册
- 用户登录并返回 JWT
- 产品创建、修改、删除、列表
- 设备创建、修改、删除、列表
- 开放平台签名验证
- 开放平台下发消息到设备
- device-rpc 通过 MQTT publish 设备命令
- device-consumer 订阅设备心跳/状态 topic 并更新数据库
- device-consumer 发布设备状态事件到 Redis Pub/Sub
- push-api 通过 SSE 给后台推送设备状态事件

## 依赖服务

当前配置对应你的 Docker 端口：

```text
MySQL  127.0.0.1:3307
Redis  127.0.0.1:16379
Etcd   127.0.0.1:12379
MQTT   127.0.0.1:1883
EMQX Dashboard http://127.0.0.1:18083
```

## 数据库

先创建数据库：

```sql
CREATE DATABASE iot DEFAULT CHARACTER SET utf8mb4;
```

然后执行 SQL：

```text
app/user/rpc/user.sql
app/product/rpc/sql/product.sql
app/device/rpc/sql/device.sql
```

## 服务端口

```text
user-rpc       8080
product-rpc    8081
device-rpc     8082

user-api       8888
product-api    8889
device-api     8890
openapi-api    8891
push-api       8892
```

## 启动顺序

先确保 MySQL、Redis、Etcd、EMQX 都已启动。

启动 RPC：

```bash
go run ./app/user/rpc/user.go -f ./app/user/rpc/etc/user.yaml
go run ./app/product/rpc/product.go -f ./app/product/rpc/etc/product.yaml
go run ./app/device/rpc/device.go -f ./app/device/rpc/etc/device.yaml
```

启动 API：

```bash
go run ./app/user/api/user.go -f ./app/user/api/etc/user-api.yaml
go run ./app/product/api/product.go -f ./app/product/api/etc/productapi-api.yaml
go run ./app/device/api/device.go -f ./app/device/api/etc/device-api.yaml
go run ./app/openapi/api/openapi.go -f ./app/openapi/api/etc/openapi-api.yaml
go run ./app/push/api/push.go -f ./app/push/api/etc/push-api.yaml
```

启动 MQTT consumer：

```bash
go run ./app/device/consumer/deviceconsumer.go -f ./app/device/consumer/etc/device-consumer.yaml
```

## MQTT Topic 约定

后端下发命令：

```text
device/{deviceKey}/command
```

设备上报心跳：

```text
device/{deviceKey}/heartbeat
```

设备上报状态：

```text
device/{deviceKey}/status
```

状态 payload：

```text
online
offline
```

## 常用接口

用户注册：

```bash
curl -X POST http://127.0.0.1:8888/api/user/register \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

用户登录：

```bash
curl -X POST http://127.0.0.1:8888/api/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

创建设备：

```bash
curl -X POST http://127.0.0.1:8890/api/device/create \
  -H "Content-Type: application/json" \
  -d '{"product_id":1,"name":"demo-device","device_key":"dev001","status":0}'
```

订阅后台设备状态 SSE：

```bash
curl http://127.0.0.1:8892/api/push/device/status
```

## 开放平台签名

开放平台接口：

```text
POST http://127.0.0.1:8891/openapi/device/send
```

签名 Header：

```text
X-App-Id
X-Timestamp
X-Nonce
X-Sign
```

当前测试配置：

```yaml
SignAuth:
  ExpireSeconds: 300
  Apps:
    test-app: test-secret
```

签名规则：

```text
sha256(appId=...&body=...&nonce=...&timestamp=...&secret=...)
```

其中参数按 key 字典序排序，`body` 是原始请求体字符串。

## 编译检查

```bash
go test ./app/user/... ./app/product/... ./app/device/... ./app/openapi/... ./app/push/... -run TestDoesNotExist -count=0
```

## go-zero 生成后要做什么

更详细的学习笔记在：

```text
app/instruction.md
```

简单记忆：

```text
RPC: proto -> goctl -> config -> yaml -> svc 注入 model/client -> logic
API: api -> goctl -> config -> yaml -> svc 注入 rpc client -> logic
```

## 后续可增强

- 用户密码改为 bcrypt/argon2 哈希存储
- 统一错误码和响应格式
- openapi nonce 防重放从内存改为 Redis
- openapi appId/secret 改为数据库管理
- MQTT ACL、设备认证和证书接入
- EMQX Webhook/Rule Engine 接入上下线事件
- 完整单元测试和集成测试
- 管理后台 API 和前端页面
