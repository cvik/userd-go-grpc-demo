package usersv1

//go:generate protoc --go_out=paths=source_relative:../pkg/services users/v1/users.proto
//go:generate protoc --go-grpc_out=paths=source_relative:../pkg/services users/v1/users.proto
