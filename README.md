# douyinlive-proto

抖音直播 protobuf 定义与 Go 生成代码。

This repository contains Douyin live protobuf definitions and generated Go code.

## 包结构

```text
protobuf/new_douyin.proto
generated/new_douyin/new_douyin.pb.go
generated/struct.go
generated/messagepool.go
```

- `protobuf/new_douyin.proto`：当前使用的 protobuf 定义。
- `generated/new_douyin`：`protoc-gen-go` 生成的 Go 结构体。
- `generated/struct.go`：消息 method 到 protobuf 结构体的映射。
- `generated/messagepool.go`：消息实例池，减少高频消息解析时的临时对象分配。

## 使用方式

在业务项目中引用：

```go
import (
	"github.com/jwwsjlm/douyinlive-proto/generated"
	"github.com/jwwsjlm/douyinlive-proto/generated/new_douyin"
)
```

## 重新生成

需要先安装 `protoc` 和 `protoc-gen-go`。

```powershell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
make proto
go test ./...
```

Windows 如果没有安装 `make`，可以直接运行：

```powershell
protoc --proto_path=protobuf --go_out=. --go_opt=module=github.com/jwwsjlm/douyinlive-proto protobuf/new_douyin.proto
```

或者：

```cmd
protobuf\new_pb.cmd
```

生成输出位置：

```text
generated/new_douyin/new_douyin.pb.go
```

## 发版流程

1. 修改 `protobuf/new_douyin.proto`。
2. 运行 `make proto`。
3. 如果新增消息类型，同步更新 `generated/struct.go` 和 `generated/messagepool.go`。
4. 运行 `go test ./...`。
5. 提交并打 tag，例如 `v0.1.1`。
6. 业务仓库通过 `go get github.com/jwwsjlm/douyinlive-proto@v0.1.1` 升级。
