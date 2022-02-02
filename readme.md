GO-GRPC BOILERPLATE
1. grpc setting
    ```
    brew install protobuf

    go get -u google.golang.org/grpc
    go get google.golang.org/protobuf/reflect/protoreflect@v1.26.0
    go get google.golang.org/protobuf/runtime/protoimpl@v1.26.0

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    ```
2. compile .proto file to go file
    ```
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc_application/grpc_application.proto
    ```
