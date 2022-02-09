# gRPC - Up and Running
This repository is for practicing grpc with go language.

# Prerequisites
* Go, any one of the three latest major releases of Go.

    For installation instructions, see Go’s Getting Started guide.

* Protocol buffer compiler, protoc, version 3.

    For installation instructions, see Protocol Buffer Compiler Installation.

* Go plugins for the protocol compiler:

    1. Install the protocol compiler plugins for Go using the following commands:

    ```bash
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    ```

    2. Update your PATH so that the protoc compiler can find the plugins:

    ```bash
    $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```

# Generate gRPC code
```
$ protoc --go_out=[output file path] --go_opt=paths=source_relative \
--go-grpc_out=[output file path] --go-grpc_opt=paths=source_relative \
[proto file path]
```


# Reference
[site](https://grpc-up-and-running.github.io/)

[sample_codes](https://github.com/grpc-up-and-running/samples)

[grpc.io - official grpc site](https://grpc.io/docs/languages/go/quickstart/#prerequisites)