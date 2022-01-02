package net.jp.ytake.concat

import com.github.ytake.morse.message.combine.definition.UserAction
import net.jp.ytake.concat.config.ConfigFinder
import org.apache.spark.sql.functions._
import org.apache.spark.sql.types.{IntegerType, LongType, StringType, TimestampType}
import scalapb.spark.Implicits._
import scalapb.spark.ProtoSQL

import java.util.Properties

object StructuredStreamingRunner extends App {

  if (args.length < 1) {
    throw new IllegalArgumentException(
      "This program takes one argument: the path to an environment configuration file.")
  }
  val ss = new SparkApplication(getClass.getName, "/tmp/streaming_runner").createSession()

  val prop = new Properties
  prop.load(new java.io.FileInputStream(args(0)))
  val df = KafkaDataFrame.make(ss, new ConfigFinder(prop))

  /**
   * root
   * |-- key: binary (nullable = true)
   * |-- value: binary (nullable = true)
   * |-- topic: string (nullable = true)
   * |-- partition: integer (nullable = true)
   * |-- offset: long (nullable = true)
   * |-- timestamp: timestamp (nullable = true)
   * |-- timestampType: integer (nullable = true)
   */
  // for udf
  val parseUserAction = ProtoSQL.udf { bytes: Array[Byte] => UserAction.parseFrom(bytes) }

  // create view "events"
  /**
   * origin
   * root
   * |-- key: string (nullable = true)
   * |-- value: struct (nullable = false)
   * |    |-- correlationId: long (nullable = true)
   * |    |-- event: string (nullable = true)
   * |    |-- userId: integer (nullable = true)
   * |    |-- name: string (nullable = true)
   * |    |-- created: struct (nullable = true)
   * |    |    |-- seconds: long (nullable = true)
   * |    |    |-- nanos: integer (nullable = true)
   * |-- correlation_id: long (nullable = true)
   * |-- event: string (nullable = true)
   * |-- user_id: integer (nullable = true)
   * |-- name: string (nullable = true)
   * |-- created: long (nullable = true)
   * |-- created_timestamp: timestamp (nullable = true)
   */
  df
    .select(col("key"), col("value"))
    .withColumn("key", col("key").cast(StringType))
    .withColumn("value", parseUserAction(col("value")))
    .withColumn("correlation_id", col("value.correlationId").cast(LongType))
    .withColumn("event", col("value.event").cast(StringType))
    .withColumn("user_id", col("value.userId").cast(IntegerType))
    .withColumn("name", col("value.name").cast(StringType))
    .withColumn("created", col("value.created.seconds").cast(LongType))
    .withColumn("created_timestamp", col("value.created.seconds").cast(TimestampType))
    // for aggregation / window
    // .withWatermark("created_timestamp", "10 minutes")
    .drop("value")
    .createTempView("events")

  /**
   * root
   * |-- key: string (nullable = true)
   * |-- correlation_id: long (nullable = true)
   * |-- event: string (nullable = true)
   * |-- user_id: integer (nullable = true)
   * |-- name: string (nullable = true)
   * |-- created: long (nullable = true)
   * |-- created_timestamp: timestamp (nullable = true)
   */
  ss
    .sql("SELECT * FROM events")
    .writeStream
    .format("console")
    .start()
    .awaitTermination()
}
