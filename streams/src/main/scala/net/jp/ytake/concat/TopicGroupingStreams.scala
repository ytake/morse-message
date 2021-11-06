package net.jp.ytake.concat

object TopicGroupingStreams extends App {
  if (args.length < 1) {
    throw new IllegalArgumentException(
      "This program takes one argument: the path to an environment configuration file.")
  }
}
