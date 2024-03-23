# protobuf example: generate go code by protobuf file

[doc](https://developers.google.com/protocol-buffers/docs/gotutorial)

download protoc <https://github.com/protocolbuffers/protobuf/releases>

```Go
go install -u github.com/golang/protobuf/protoc-gen-go@latest
go generate
go build
```