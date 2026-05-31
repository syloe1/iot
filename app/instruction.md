# go-zero Module Checklist

This project uses a common go-zero monorepo flow:

```text
api -> rpc -> model -> mysql
```

## RPC module after proto generation

After writing a `.proto` file and running `goctl rpc protoc`, do these steps:

1. Check generated files.
   - `xxx.go` is the rpc entrypoint.
   - `etc/xxx.yaml` is runtime config.
   - `internal/config/config.go` defines yaml fields.
   - `internal/svc/servicecontext.go` wires dependencies.
   - `internal/logic/*.go` contains business code.
   - `xxxclient` or `xxxsvc` is the rpc client package used by api modules.

2. Add config fields in `internal/config/config.go`.
   - If rpc needs MySQL, add `Mysql.DataSource`.
   - If generated model uses cache, add `CacheRedis cache.CacheConf`.
   - Add other external dependencies here, such as JWT config or Redis config.

3. Fill `etc/xxx.yaml`.
   - Set `Name` and `ListenOn`.
   - Add `Mysql.DataSource` if the service uses a database.
   - Add `CacheRedis` if the model is cached.
   - Use either `Etcd` service discovery or direct ports consistently.

4. Wire dependencies in `internal/svc/servicecontext.go`.
   - Create models with `sqlx.NewMysql(c.Mysql.DataSource)`.
   - Put shared dependencies on `ServiceContext`.
   - Logic files should use `l.svcCtx.Xxx`, not create clients or db connections by themselves.

5. Implement `internal/logic/*.go`.
   - Validate request params.
   - Call model methods.
   - Convert model data into proto response data.
   - Return errors instead of swallowing them.

6. Extend model only in non-generated files.
   - Do not edit `*_gen.go`.
   - Add custom methods to `xxxmodel.go`.
   - Add custom methods to the model interface before implementing them.

7. Verify.
   - Run `go test ./app/<module>/rpc/... -run TestDoesNotExist -count=0`.
   - Start with `go run ./app/<module>/rpc/<module>.go -f ./app/<module>/rpc/etc/<module>.yaml`.

## API module after api generation

After writing a `.api` file and running `goctl api go`, do these steps:

1. Check generated files.
   - The api entrypoint should sit beside `etc` and `internal`.
   - `internal/handler` parses HTTP requests.
   - `internal/logic` is where HTTP business orchestration lives.
   - `internal/types` contains request and response structs.
   - `internal/svc/servicecontext.go` wires rpc clients.

2. Keep the output directory consistent.
   - If you want `app/product/api/internal`, generate with `-dir ./app/product/api`.
   - If you use `-dir ./app/product/api/productapi`, goctl creates a nested module-like folder.
   - Mixed output causes confusing layouts where entrypoint and internal packages are in different places.

3. Add rpc client config in `internal/config/config.go`.
   - Example:
     ```go
     type Config struct {
     	rest.RestConf
     	ProductRpc zrpc.RpcClientConf
     }
     ```

4. Fill `etc/xxx-api.yaml`.
   - Set `Name`, `Host`, and `Port`.
   - For beginner learning, direct endpoints are easier:
     ```yaml
     ProductRpc:
       Endpoints:
         - 127.0.0.1:8081
     ```
   - If using Etcd, the api client `Etcd.Hosts` must match the rpc server `Etcd.Hosts`.

5. Wire rpc clients in `internal/svc/servicecontext.go`.
   - Example:
     ```go
     ProductRpc: productsvc.NewProductSvc(zrpc.MustNewClient(c.ProductRpc))
     ```
   - API logic should call rpc clients through `svcCtx`.

6. Implement `internal/logic/*.go`.
   - Convert api request structs to rpc request structs.
   - Call rpc methods.
   - Convert rpc response structs to api response structs.

7. Be careful with JWT.
   - Login routes should not be under `@server(jwt: JwtAuth)`.
   - Protected routes can use a separate `@server` block with JWT.

8. Verify.
   - Run `go test ./app/<module>/api/... -run TestDoesNotExist -count=0`.
   - Start rpc first, then api.
   - Test with curl or an HTTP client.

## Directory examples in this repo

Preferred layout:

```text
app/user/api/user.go
app/user/api/etc
app/user/api/internal
app/user/api/userapi.api
```

Use the same style for product:

```text
app/product/api/product.go
app/product/api/etc
app/product/api/internal
app/product/api/productapi.api
```

If a nested folder such as `app/product/api/productapi` appears, it usually means `goctl api go` was run with a nested `-dir`. Pick one layout and keep using the same `-dir`.

## MQTT device messaging flow

The device side now uses MQTT instead of a hand-written WebSocket connector.

Command flow:

```text
openapi-api -> device-rpc -> MQTT publish -> device
```

Status flow:

```text
device -> MQTT publish -> device-consumer -> mysql + Redis Pub/Sub -> push-api SSE -> admin frontend
```

Topic conventions:

```text
device/{deviceKey}/command    backend publishes command payloads here
device/{deviceKey}/heartbeat  device publishes heartbeat here
device/{deviceKey}/status     device publishes "online" or "offline" here
```

Runtime dependencies:

```text
MySQL  127.0.0.1:3307
Etcd   127.0.0.1:12379
MQTT   127.0.0.1:1883
Redis  127.0.0.1:16379
```

Service startup example:

```bash
go run ./app/device/rpc/device.go -f ./app/device/rpc/etc/device.yaml
go run ./app/device/consumer/deviceconsumer.go -f ./app/device/consumer/etc/device-consumer.yaml
go run ./app/openapi/api/openapi.go -f ./app/openapi/api/etc/openapi-api.yaml
go run ./app/push/api/push.go -f ./app/push/api/etc/push-api.yaml
```

Admin frontend can subscribe to device status events with:

```text
GET http://127.0.0.1:8892/api/push/device/status
```
