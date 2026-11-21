# Noah API 定义与 Swagger

这里的 `.proto` 是接口定义源文件；不要直接修改生成的 Swagger。

```sh
# 仓库根目录；需 buf 和 buf.gen.yaml 中的四个本地 protoc 插件
make buf
# 只检查 proto 是否能编译
cd apis && buf build
```

`make buf` 生成 `gen/go/` 和 **`output/apis.swagger.yaml`**（Swagger 2.0）。
这两个目录已被 Git 忽略。生成器使用 `noah/gen/go` 包前缀。
Swagger 未固定 host/scheme；Swagger UI 应与 API 同源，或自行配置目标服务器。

## 可以调用的 HTTP 接口

以下路径以 `apps/*` 中注册的直接 HTTP handler 为准。请求字段使用
**snake_case**：这些 handler 使用 Go `encoding/json`，不是 grpc-gateway 的 protojson。
POST 请求除 OpenAI 代理以外，业务字段直接放 JSON body，不要包装成 `payload`。
GET 响应保持 handler 的实际字段大小写：部分列表直接输出 Go/GORM 字段名，
例如 `ID`、`LoginName`、`CreatedAt`；公司业务字段则使用 `name`、`uen`。
示例地址中的 `localhost:8080` 请替换为实际服务地址。

| POST 路径 | JSON body 示例 | 当前成功响应 |
| --- | --- | --- |
| `/apis/infra/v1/host/heartbeat` | `{"ip":"203.0.113.10","hostname":"web-01","arch":"arm64","os":"linux"}` | `{}` |
| `/apis/infra/v1/linux/{hostname}` | `{"action":"login"}` | `record visit success`（文本） |
| `/apis/v1/company/create` | `{"name":"Example Pte Ltd","uen":"202400001A","website":"https://example.com","status":1}` | 空 body |
| `/apis/v1/sgw/sec/check` | `{"user_agent":"Mozilla/5.0","method":"GET","url":"https://example.com/login","request_id":"req-example-001"}` | `record success`（文本） |
| `/v1/sensitive/create` | `{"category":"spam","country_code":"SG","content":"example-word"}` | `create sensitive`（文本） |
| `/v1/ipdb/clientip/create` | `{"ip":"203.0.113.10","isp_name":"Example ISP","cnt":"1"}` | `create client ip success`（文本） |
| `/v1/ipdb/datacenter/create` | `{"name":"example-datacenter","domain":"example.net","ip_start":"203.0.113.0","ip_end":"203.0.113.255","ip_start_int":3405803776,"ip_end_int":3405804031}` | `create datacenter`（文本） |
| `/v1/openai/{name}` | 见下方代理示例；`name` 支持多段 upstream 路径 | upstream JSON |

```sh
curl -X POST http://localhost:8080/apis/infra/v1/host/heartbeat \
  -H 'Content-Type: application/json' \
  -d '{"ip":"203.0.113.10","hostname":"web-01","arch":"arm64","os":"linux"}'

curl -X POST http://localhost:8080/v1/openai/v1/chat/completions \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <token>' \
  -d '{"model":"your-model","messages":[{"role":"user","content":"Hello"}],"stream":false}'
```

代理的 body 是原始 upstream JSON。代理 payload 使用 `google.protobuf.Struct`，
不需要 `google.protobuf.Any` 的 `@type` 包装。Swagger UI 如果把包含 `/` 的
path 参数编码为 `%2F`，请使用上面的完整 URL 调用。

## 实现边界

- 标记 `deprecated` 的 Azure、UA、UA OS 和旧 UserKey 操作只有 proto 合同，当前没有对应注册路由，不能据此直接调用。
- 安全检查当前只入库，不返回 allow/block 决策；敏感词接口用于新增词条，不执行文本扫描。
- `payload` / `data` 是多数旧消息保留的扩展字段，直接 HTTP 示例应省略。
- 心跳的 `received_at` / `config_updates` 当前不赋值；公司创建成功不返回完整 profile。
- 错误处理尚未统一：有些 handler 返回文本，有些在 HTTP 200 中携带错误状态，有些先写 body 再设置状态。生成器自动附带的 `default` gRPC 错误模型只适用于 gateway，不能视为所有直接 handler 的错误合同。
- OpenAI handler 当前没有完整转发 upstream HTTP 状态和 headers，也没有真正的流式转发；示例使用 `stream:false`。
- 当前 GET 文档覆盖业务读取、健康检查、基础查询及下载接口；不包括兜底 404、pprof 标准调试中间件及 SGW 根路径欢迎页。上传和部分 POST 管理接口仍需另行补充。新增接口时应同步注册路由和 proto 合同。

## 注释和示例规范

1. service tag 说明真实用途；RPC 的 `openapiv2_operation` 提供唯一 operation ID、摘要、行为及实现限制。
2. `google.api.http` 的路径、body 和 response_body 必须与 handler 一致；完整 JSON body 使用 `body: "*"`。
3. 字段说明及 JSON 示例放在 `openapiv2_field`；完整请求示例放在 `openapiv2_schema`，避免只写不会驱动 Swagger 示例的普通注释。
4. 示例必须与字段类型一致；int64 在 protobuf JSON 中用字符串表达。必填标记与验证约束需有实现依据。
5. 修改后执行 `buf format -w`、`buf build`、`buf generate`，检查生成的引用、operation ID、path 参数、body 和响应类型。可联网时再执行 `buf lint`；现有命名规则问题需要单独处理。

## GET 接口

本轮补齐 37 个 GET 操作（包含别名），都不需要请求 body。具体响应示例在
Swagger 的 `responses -> 200 -> examples` 中；列表是直接 JSON 数组，不带包装对象。
列表目前固定返回前 10 条，company 与 SGW 为前 50 条，未实现 `limit` / `offset` 查询参数。

| 分组 | GET 路径 |
| --- | --- |
| 用户和密钥 | `/v1/user`、`/v1/user/list`、`/v1/user/keys`、`/v1/user/key`、`/v1/user/key/list` |
| 敏感词 | `/v1/sensitive`、`/v1/sensitive/list` |
| IP 数据 | `/v1/ipdb/datacenter`、`/v1/ipdb/datacenter/list`、`/v1/ipdb/clientip`、`/v1/ipdb/clientip/list` |
| 公司 | `/apis/v1/company`（也支持尾部 `/`） |
| 主机 | `/apis/infra`、`/apis/infra/v1`、`/apis/infra/v1/linux/{hostname}`、`/apis/infra/v1/motd/{hostname}`、`/apis/infra/v1/linux/count` |
| SGW | `/apis/v1/sgw/domain`、`/apis/v1/sgw/dns`、`/apis/v1/sgw/cert`、`/apis/v1/sgw/alb` |
| 健康检查 | `/healthz`、`/apis/v1/sgw/sec/check`、`/internal/security/check`、`/internal/metrics` |
| 基础查询 | `/`、`/help`、`/apis`、`/echo`、`/info`、`/json`、`/{ip}/json` |
| 下载 | `/downloads/{filename}`、`/dl/{filename}`（建议使用前者，后者存在已有的前缀截取问题） |
| 用户服务 | `/apis/users`、`/apis/users/oauth2/v1/callback`、`/apis/oauth2/v1/callback` |

```sh
curl http://localhost:8080/v1/user/list
curl http://localhost:8080/apis/v1/company
curl http://localhost:8080/apis/v1/sgw/sec/check
curl 'http://localhost:8080/apis/infra/v1/linux/web-01?action=login'
```

请按当前实现理解这些接口：

- `/v1/user`、`/v1/user/key`、两个 IP 单条接口只返回占位文本；敏感词单条接口固定读取 ID=1。
- SGW 的 domain/dns/cert/alb 当前共用 `ListDomain`，均返回域名列表。
- 主机 visit 的 GET 会写入记录，`action` 是实际支持的 query 参数。
- OAuth callback 当前只是诊断回显，不执行认证；`/internal/metrics` 只是返回 `ok`。
- `ListUsers` 和 `ListKeys` 当前直接序列化数据库模型，包含 `Password` / `Token` 字段。
  文档示例使用虚构或脱敏值，本轮未改变接口的实际返回行为。
- 通用的 `TextResponse` / `ObjectResponse` / `ListResponse` 只用于描述直接 HTTP 响应，通过 `response_body: "data"`
  表达文本、对象或数组；不会给现有 handler 增加 `data` 外层，也不会自动实现 gRPC 服务。

## 文件组织

按业务资源组织接口，同一个资源的 GET、POST 及后续更新/删除方法放在同一个
proto 文件和 service 中，不按读写或 HTTP 方法拆分。

| 业务资源 | 定义文件 |
| --- | --- |
| 用户资料和 SSH 公钥 | `user/v1/profile.proto` |
| 用户 API key | `user/v1/keys.proto` |
| 用户认证回调 | `user/v1/auth.proto` |
| 敏感词 | `apigw/v1/sensitive.proto` |
| 公司资料 | `company/v1/profile.proto` |
| 客户端 IP / 数据中心 | `ipinfo/v1/client_ip.proto` / `ipinfo/v1/datacenter.proto` |
| 主机访问记录和统计 | `infra/v1/motd.proto` |
| 安全检查与健康检查 | `security/v1/sec_check.proto` |
| 网关资源 | `sgw/v1/gateway.proto` |
| 基础服务、诊断和下载 | `system/v1/service.proto` |

`common/v1/http.proto` 仅保留可复用的 HTTP 响应类型，不定义独立读服务。
