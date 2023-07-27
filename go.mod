module github.com/MikeMwita/fedha.git

go 1.21rc2

replace github.com/MikeMwita/fedha-go-gen.grpc => /home/mike/Desktop/fedha-go-gen.grpc

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230706204954-ccb25ca9f130 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230724170836-66ad5b6ff146 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230724170836-66ad5b6ff146 // indirect
	google.golang.org/grpc v1.56.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
