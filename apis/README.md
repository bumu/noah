# Noah API 定义目录

本目录包含Noah系统的所有API接口定义。

## 目录结构

```
apis/
├── internal/     # 内部API定义
├── external/     # 外部/公开API定义
└── common/       # 通用API定义
```

## 使用说明

1. 所有API定义文件应遵循OpenAPI/Swagger规范
2. 文件命名规则：`<service-name>_api.yaml`
3. 每个API定义文件需包含:
   - 接口描述
   - 请求/响应格式
   - 错误码说明
   - 版本信息

## 开发规范

- API版本控制采用语义化版本(Semantic Versioning)
- 所有API需包含适当的认证和授权机制
- 确保API文档及时更新
