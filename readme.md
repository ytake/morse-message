# メッセージを分割しすぎると大変なことになるサンプル

WIP

## Required 

### Protocol Buffer Compiler Installation

```bash
$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
```

or

```bash
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

Compiler Invocation (Go)

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```bash
$ make gen
```
