# Go Publisher / Subscriber

動かすための事前準備

## for Mac(and M1)

```bash
$ brew install openssl
$ brew install librdkafka
$ brew install pkg-config
# 実際の環境に合わせてください(以下はM1 Mac向け)
$ export PKG_CONFIG_PATH="/opt/homebrew/opt/openssl@3/lib/pkgconfig"
$ go run --tags dynamic main.go
```

## Env 

```bash
$ export KAFKA_BOOTSTRAP_SERVERS=127.0.0.1:29092
```

## Run 

### topic single-user-action 

*1 Topic / 1 Partition 構成*  
シンプルな構成はなにもしなくてもうまくいく（ように見える）という例

以下のコマンドでメッセージを送信します

```bash
$ go run --tags dynamic main.go m:spp
```

以下のコマンドでメッセージを取得します

```bash
$ go run --tags dynamic main.go m:sps
```

### topic nokey-user-action

*1 Topic / 2(n) Partition 構成*  

この例ではuserの登録削除動作をバラバラに分散するため、  
想定しない時系列を生んでしまうことなります。  
実際に組み込む場合にはこのパターンにならないように気をつけなければならない、という例です。

以下のコマンドでメッセージを送信します

```bash
$ go run --tags dynamic main.go m:nkpp
```

以下のコマンドでメッセージを取得します

```bash
$ go run --tags dynamic main.go m:nkps
```

Consumerが1プロセスの場合は、1プロセスで全Partitionから取得します  
取得時は下記のように標準出力されます。  
`[]` 内の数字は取得元のPartitionを表します。

```bash
partition nokey-user-action[1]@0Event: CREATED
```

### topic haskey-user-action

*1 Topic / 2(n) Partition 構成*

この例ではuserの登録削除動作を分散方法を指示し、  
想定しない時系列を生まないものとなっています。  

以下のコマンドでメッセージを送信します

```bash
$ go run --tags dynamic main.go m:hkpp
```

以下のコマンドでメッセージを取得します

```bash
$ go run --tags dynamic main.go m:hkps
```

Consumerが1プロセスの場合は、1プロセスで全Partitionから取得します    
取得時は下記のように標準出力されます。  
`[]` 内の数字は取得元のPartitionを表します。

```bash
partition nokey-user-action[1]@0Event: CREATED
```
