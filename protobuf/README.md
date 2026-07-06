# Protobuf 生成说明

本文档说明如何从 `protobuf/new_douyin.proto` 重新生成 Go 代码。

## 目录说明

- `new_douyin.proto`：当前 DouyinLive 使用的 protobuf 定义。
- `webcast.data.proto` 和 `webcast/data/*.proto`：历史或参考定义，默认生成流程不会使用。
- `../generated/new_douyin/new_douyin.pb.go`：由 `new_douyin.proto` 生成的 Go 代码。
- `../generated/messagepool.go` 和 `../generated/struct.go`：项目维护的消息池和 method 到消息类型的映射，不由 `protoc` 自动生成。

## 环境要求

需要安装两个工具：

- `protoc`：Protocol Buffers 编译器。
- `protoc-gen-go`：Go 代码生成插件。

本仓库的 CI 工具链策略：

- `protoc`：固定为 `35.1`，避免生成文件头部的编译器版本号频繁变化。
- `protoc-gen-go`：使用 `latest`，跟随 `google.golang.org/protobuf` 最新稳定版本。

建议本地生成时使用同样的 `protoc` 版本，并安装最新的 `protoc-gen-go`，否则 `generated/new_douyin/new_douyin.pb.go` 头部记录的工具版本不同，CI 会认为生成文件没有同步提交。

安装 `protoc-gen-go`：

```powershell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

确认 `protoc-gen-go` 已经在 `PATH` 中：

```powershell
protoc-gen-go --version
```

确认 `protoc` 可用：

```powershell
protoc --version
```

## 生成命令

在仓库根目录执行：

```powershell
protoc --proto_path=protobuf --go_out=. --go_opt=module=github.com/jwwsjlm/douyinlive-proto protobuf/new_douyin.proto
```

如果有 `make`：

```powershell
make proto
```

Windows 可以运行：

```cmd
protobuf\new_pb.cmd
```

生成后的文件位置：

```text
generated/new_douyin/new_douyin.pb.go
```

## 重要的 go_package 配置

`new_douyin.proto` 必须保留这个配置：

```proto
option go_package = "github.com/jwwsjlm/douyinlive-proto/generated/new_douyin;new_douyin";
```

生成时必须带上：

```text
--go_opt=module=github.com/jwwsjlm/douyinlive-proto
```

这个参数会裁剪 module 路径，让生成文件保持在 `generated/new_douyin/` 目录下。

## 修改 proto 后的流程

1. 修改 `protobuf/new_douyin.proto`。
2. 重新生成 Go 代码。
3. 如果消息 method 名称有变化，更新 `generated/struct.go`。
4. 如果消息池逻辑需要变化，更新 `generated/messagepool.go`。
5. 运行检查：

```powershell
go test ./...
go vet ./...
go run honnef.co/go/tools/cmd/staticcheck@latest ./...
protoc --proto_path=protobuf --go_out=. --go_opt=module=github.com/jwwsjlm/douyinlive-proto protobuf/new_douyin.proto
git diff --exit-code -- generated/new_douyin/new_douyin.pb.go
```

如果最后一步只显示 `protoc` 版本号差异，先检查本地 `protoc --version` 是否和 CI 固定版本一致。

## 注意事项

- 不要手动编辑 `generated/new_douyin/new_douyin.pb.go`，应该修改 `.proto` 后重新生成。
- 不要把旧的 `webcast.data.proto` 路径无脑合并到默认生成流程，除非明确要迁移协议来源。
- 修改 protobuf 定义时，要把生成后的 Go 文件一起提交。
