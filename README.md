### Userd - A simple demo service using gRPC in Go

#### Required Tooling

You need to install the protobuf compiler (`protoc`).
See https://grpc.io/docs/protoc-installation/ for instructions for your OS.

Then install the Go plugins:

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

#### To Build

```bash
$ go generate ./... # Should do nothing since the generated code is checked in
$ go run cmd/server/main.go
```

#### How to test

Install `grpcurl`:

```bash
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7`
```

Then run the server in a separate shell `go run cmd/servermain.go`.

When the server is running we can call it with grpcurl like so:

```bash
$ grpcurl -plaintext -d '{"name": "Leffe", "Email": "leffe@awesome.io"}' localhost:1998 lendo.users.v1.Users/CreateUser
$ grpcurl -plaintext -d '{"id": "e7390ac4-cec8-4a54-bfc1-b3c322b61128"}' localhost:1998 lendo.users.v1.Users/GetUser
$ grpcurl -plaintext -d '{"id": "e7390ac4-cec8-4a54-bfc1-b3c322b61128"}' localhost:1998 lendo.users.v1.Users/DeleteUser
$ grpcurl -plaintext localhost:1998 lendo.users.v1.Users/ListUsers
```

You can inspect what is supported by the service by using it's built-in reflection
capabilities:

```bash
$ grpcurl -plaintext localhost:1998 list
$ grpcurl -plaintext localhost:1998 describe lendo.users.v1.Users
```

You can describe individual RPCs and messages as well. See `grpcurl --help` for more info.
