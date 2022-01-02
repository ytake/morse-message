package net.jp.ytake.concat

import net.jp.ytake.concat.config.ConfigFinder
import org.apache.spark.sql.{DataFrame, SparkSession}

object KafkaDataFrame {

  /**
   * return DataFrame
   * @param ss
   * @param cf
   * @return
   */
  def make(ss: SparkSession, cf: ConfigFinder): DataFrame = {
    ss.readStream
      .format("kafka")
      .option("kafka.bootstrap.servers", cf.bootstrapServers)
      .option("subscribe", cf.topic)
      .option("startingOffsets", "earliest")
      .load()
  }
}
