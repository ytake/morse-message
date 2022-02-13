# 適切な構成とPartitionを用いないとメッセージがうまく扱えない集

このリポジトリは

 - 1Partitionの場合は簡単に扱える
 - 2Partition以上の場合に、分散方法を指定しないとどうなるか
 - バラバラに分散されたものはSparkなどでストリーム処理を挟んでも完全には戻せない

という、例を簡単に確認できるようにしたもの  

各アプリケーションの使い方  

 - [Producr / Consume(Go)](https://github.com/ytake/morse-message/tree/main/go)
 - [Structured Streaming(Scala)](https://github.com/ytake/morse-message/tree/main/streams)

メッセージにはProtocol Bufferを利用しています  

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
$ docker exec kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --if-not-exists --bootstrap-server kafka:9092 --replication-factor 1 --partitions 2 --topic nokey-user-action
$ docker exec kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --if-not-exists --bootstrap-server kafka:9092 --replication-factor 1 --partitions 2 --topic haskey-user-action
$ docker exec kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --if-not-exists --bootstrap-server kafka:9092 --replication-factor 1 --partitions 1 --topic single-user-action
```

delete topics 例

```bash
$ docker-compose exec broker pub-topics --bootstrap-server kafka:9092 --delete --topic user-action-created
```
