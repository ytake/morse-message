package net.jp.ytake.concat

import org.apache.spark.sql.SparkSession
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @param appName
 * @param checkPoint
 */
class SparkApplication(appName: String, checkPoint: String) {

  protected def context(): SparkContext = {
    val conf = new SparkConf().setAppName(appName)
    conf.set("spark.speculation", "false")
    conf.set("spark.sql.session.timeZone", "Asia/Tokyo")
    new SparkContext(conf)
  }

  def createSession(): SparkSession = {
    val spark = context()
    spark.hadoopConfiguration.set("parquet.enable.dictionary", "false")
    spark.hadoopConfiguration.set("mapreduce.fileoutputcommitter.algorithm.version", "2")
    spark.hadoopConfiguration.set("mapreduce.fileoutputcommitter.cleanup-failures.ignored", "true")
    spark.setCheckpointDir(checkPoint)
    spark.setLogLevel("WARN")
    SparkSession.builder
      .appName(appName)
      .getOrCreate
  }
}
