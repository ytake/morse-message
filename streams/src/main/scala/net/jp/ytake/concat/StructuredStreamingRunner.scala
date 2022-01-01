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
    .drop("value")
    .createTempView("events")

  /**
   * table example
   * +----+--------------+-------+-------+------+----------+-------------------+
   * | key|correlation_id|  event|user_id|  name|   created|  created_timestamp|
   * +----+--------------+-------+-------+------+----------+-------------------+
   * |null|             0|CREATED|      1|  aaa1|1598996803|2020-09-02 06:46:43|
   * |null|             0|DELETED|      1|  aaa1|1598996803|2020-09-02 06:46:43|
   * |null|             0|CREATED|      3|  aaa3|1598996803|2020-09-02 06:46:43|
   * +----+--------------+-------+-------+------+----------+-------------------+
   */
  ss
    .sql("SELECT * FROM events")
    .writeStream
    .format("console")
    .start()
    .awaitTermination()
}
