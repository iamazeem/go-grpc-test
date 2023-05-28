# go-grpc-test

Go [GRPC](https://grpc.io/) Test

- sample DNS client and server apps

## Set up

- Install protocol buffers compiler (`protoc`)

- Install protocol buffers and GRPC plugins:

  ```shell
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  ```

  Update `GOPATH`:

  ```shell
  export PATH="$PATH:$(go env GOPATH)/bin"
  ```

- Generate protocol buffers and GRPC files:

  ```shell
  protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    dns/dns.proto
  ```

  Following files are generated:

  - `dns/dns.pb.go`
  - `dns_grpc.pb.go`

  The generated files are already under `dns` directory.

## Build

```shell
go mod download
go build -o dns-server server/main.go
go build -o dns-client client/main.go
```

## Run

Run server in a separate terminal (as it blocks):

```shell
$ ./dns-server
staring DNS GRPC server on 0.0.0.0:9000
gRPC server started listening on 0.0.0.0:9000
```

Run client from a different terminal:

```shell
$ ./dns-client --domain "domain.xyz"
192.168.0.0
```

## License

[MIT](./LICENSE)
