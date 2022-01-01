# Structured Streaming + Kafka

Sparkを利用して、Kafkaトピックのレコードを取得する例  

マイクロバッチでの操作になるため、過去のデータを扱うには  
通常の読み込みを利用してunionなどをする必要がありますが、  
メッセージのタイムラインを分轄するようなレコードになってしまうと（結果的になってしまった場合など）  
時系列をうまく作り直すのは難しくなります。

特殊な用途がない限り、メッセージの分割は行わずにメインのタイムラインは維持しながら  
groupなどでアプリケーションに合わせて処理しましょう。

## Compile

```bash
$ sbt assembly
```

## Spark Submit

```bash
$ spark-submit target/scala-2.12/split-brain-groups-1.0.jar ./config/kafka.properties --master spark://localhost:7077
```
