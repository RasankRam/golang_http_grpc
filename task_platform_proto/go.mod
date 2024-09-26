module example.com/task_platform_proto

go 1.23.0

replace example.com/task_platform_proto => ../task_platform_proto

require (
	github.com/envoyproxy/protoc-gen-validate v1.1.0
	google.golang.org/grpc v1.67.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
)
