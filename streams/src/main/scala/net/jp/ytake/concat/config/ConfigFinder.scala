package net.jp.ytake.concat.config

import java.util.Properties

class ConfigFinder(private val c: Properties) {
  
  def bootstrapServers: String = c.getProperty("bootstrap_servers")

  def topic: String = c.getProperty("topic")
}
