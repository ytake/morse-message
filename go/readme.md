# Go Publisher / Subscriber

動かすための事前準備

## for Mac

```bash
brew install openssl
brew install librdkafka
brew install pkg-config
# 実際の環境に合わせてください
export PKG_CONFIG_PATH="/opt/homebrew/opt/openssl@3/lib/pkgconfig"
go build --tags dynamic main.go
```

## Env 

```bash
$ export KAFKA_BOOTSTRAP_SERVERS=127.0.0.1:9092
```

## Run 

### example1 

publish messages

```bash
$ go run --tags dynamic main.go m:np
```

