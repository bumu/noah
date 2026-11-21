# Company profile API

`POST /apis/v1/company/create` 接收扁平 JSON body，`name` 必填：

```json
{
  "name": "Example Pte Ltd",
  "uen": "202400001A",
  "address": "1 Example Road, Singapore",
  "website": "https://example.com",
  "phone": "+6561234567",
  "status": 1
}
```

当前成功响应为 HTTP 200、空 body。校验或存储错误返回 `{"status":400}`，
HTTP 状态可能仍为 200。`status` 是存储字段，代码尚未定义统一的业务枚举。
完整生成命令和限制见 [API 说明](../../README.md)。
