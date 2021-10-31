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

## Create Kafka Topics

```bash
$ docker-compose exec broker pub-topics --create --zookeeper \
zookeeper:2181 --replication-factor 1 --partitions 2 --topic user-action-created
$ docker-compose exec broker pub-topics --create --zookeeper \
zookeeper:2181 --replication-factor 1 --partitions 2 --topic user-action-deleted
$ docker-compose exec broker pub-topics --create --zookeeper \
zookeeper:2181 --replication-factor 1 --partitions 2 --topic created-user-action
$ docker-compose exec broker pub-topics --create --zookeeper \
zookeeper:2181 --replication-factor 1 --partitions 1 --topic single-created-user-action
```

delete topics 

```bash
$ docker-compose exec broker pub-topics --zookeeper zookeeper:2181 --delete --topic user-action-created
```
