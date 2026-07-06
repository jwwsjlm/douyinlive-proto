# douyinlive-proto

抖音直播 protobuf 定义和 Go 生成代码。

这个仓库只维护协议层内容，业务监听、WebSocket 转发、签名逻辑不放在这里。主项目需要 protobuf 类型时，直接依赖本仓库的 Go module。

## 目录结构

```text
protobuf/new_douyin.proto
generated/new_douyin/new_douyin.pb.go
generated/struct.go
generated/messagepool.go
```

- `protobuf/new_douyin.proto`：当前 DouyinLive 使用的 protobuf 定义。
- `generated/new_douyin`：由 `protoc-gen-go` 生成的 Go 结构体。
- `generated/struct.go`：抖音 webcast method 到 protobuf 消息结构体的映射。
- `generated/messagepool.go`：消息实例池，用于减少高频解析时的临时对象分配。

## 使用方式

业务项目中这样引用：

```go
import (
	"github.com/jwwsjlm/douyinlive-proto/generated"
	"github.com/jwwsjlm/douyinlive-proto/generated/new_douyin"
)
```

升级依赖示例：

```powershell
go get github.com/jwwsjlm/douyinlive-proto@v0.1.1
go mod tidy
```

## 重新生成 Go 代码

先安装 `protoc` 和 `protoc-gen-go`：

```powershell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

如果本机有 `make`：

```powershell
make proto
go test ./...
```

Windows 没有 `make` 时，可以直接运行：

```powershell
protoc --proto_path=protobuf --go_out=. --go_opt=module=github.com/jwwsjlm/douyinlive-proto protobuf/new_douyin.proto
```

或者运行批处理：

```cmd
protobuf\new_pb.cmd
```

生成文件位置：

```text
generated/new_douyin/new_douyin.pb.go
```

## 发版流程

1. 修改 `protobuf/new_douyin.proto`。
2. 运行 `make proto` 或等价的 `protoc` 命令。
3. 如果新增了消息类型，同步更新 `generated/struct.go` 和 `generated/messagepool.go`。
4. 运行 `go test ./...`。
5. 提交并打 tag，例如 `v0.1.2`。
6. 业务仓库通过 `go get github.com/jwwsjlm/douyinlive-proto@v0.1.2` 升级。

## CI 检查

GitHub Actions 会检查：

- 重新生成 protobuf 后 `generated/new_douyin/new_douyin.pb.go` 是否和仓库一致。
- `go test ./...`
- `go vet ./...`
- `staticcheck`

这样可以避免只改了 `.proto`，但忘记提交生成文件。

## 许可证

本项目使用 MIT License，详见 [LICENSE](LICENSE)。
