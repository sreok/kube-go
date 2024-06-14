# kube-go

|                  | Kubernetes1.24 | Kubernetes1.25 | Kubernetes1.26 | Kubernetes1.27 | Kubernetes1.28 | Kubernetes1.29 | Kubernetes1.30 |
|:----------------:|:--------------:|:--------------:|:--------------:|:--------------:|:--------------:|:--------------:|:--------------:|
| `kube-go:v0.1.0` |       ✕        |       ✓        |       ✓        |       ✓        |       ✓        |       ✓        |       ✓        |


打造通用kubernetes api

- [x] 基于gin框架的后端api
- [x] 集成client-go
- [x] 集成GinSwagger
- [ ] 集成mysql、redis

## 快速开始

```shell
go mod tidy
go build -o ./build/ ./main/main.go
go run main/main.go
```

### 查看API文档（测试）
`http://127.0.0.1:8080/swagger/index.html`
### 更新API文档
```shell
swag init -g ./main/main.go -o ./docs/swagger
```