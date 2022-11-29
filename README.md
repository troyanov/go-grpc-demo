# go-grpc-demo

Generate server & client
```bash
protoc -I ./proto \
  --go_out ./proto --go_opt paths=source_relative \
  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
  ./proto/hello/hello.proto
```
