package net.jp.ytake.concat

import io.confluent.kafka.serializers.AbstractKafkaSchemaSerDeConfig
import org.apache.kafka.streams.StreamsConfig
import java.io.{FileInputStream, IOException}

object StreamProperties {

  def load(configPath: String): (EnvProperties, StreamsProperties) = {
    val envProp = loadEnvProperties(configPath)
    (envProp, buildStreamsProperties(envProp))
  }

  @throws[IOException]
  private def loadEnvProperties(fileName: String): EnvProperties = {
    val envProps = new EnvProperties()
    val input = new FileInputStream(fileName)
    envProps.load(input)
    input.close()
    envProps
  }

  private def buildStreamsProperties(envProps: EnvProperties): StreamsProperties = {
    val props = new StreamsProperties()
    props.put(StreamsConfig.APPLICATION_ID_CONFIG, envProps.getProperty("application.id"))
    props.put(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG, envProps.getProperty("bootstrap.servers"))
    props.put(StreamsConfig.DEFAULT_KEY_SERDE_CLASS_CONFIG, getClass.getName)
    props.put(StreamsConfig.DEFAULT_VALUE_SERDE_CLASS_CONFIG, getClass.getName)
    props.put(AbstractKafkaSchemaSerDeConfig.SCHEMA_REGISTRY_URL_CONFIG, envProps.getProperty("schema.registry.url"))
    props
  }
}
