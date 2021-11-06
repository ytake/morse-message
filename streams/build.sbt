ThisBuild / version := "1.0"
ThisBuild / scalaVersion := "2.13.7"
ThisBuild / organization := "net.jp.ytake"

name := "split-brain-groups"
logLevel := Level.Error
resolvers ++= Seq(
  "mvn" at "https://mvnrepository.com/artifact",
  "confluent" at "https://packages.confluent.io/maven"
)

libraryDependencies ++= Seq(
  "org.scala-lang" % "scala-library" % "2.13.7",
  "com.typesafe" % "config" % "1.2.0",
  "io.confluent" % "kafka-streams-protobuf-serde" % "6.1.3" % "provided",
  "org.apache.kafka" % "kafka-streams" % "2.7.1" % "provided",
  "org.apache.kafka" % "kafka-streams-test-utils" % "2.7.1" % Test,
  "junit" % "junit" % "4.13" % Test,
  "org.scalatest" %% "scalatest" % "3.2.7" % Test
)
assembly / assemblyMergeStrategy := {
  case m if m.toLowerCase.endsWith("manifest.mf") => MergeStrategy.discard
  case m if m.toLowerCase.matches("meta-inf.*\\.sf$") => MergeStrategy.discard
  case PathList("org","aopalliance", xs @ _*) => MergeStrategy.last
  case PathList("javax", "inject", xs @ _*) => MergeStrategy.last
  case PathList("javax", "servlet", xs @ _*) => MergeStrategy.last
  case PathList("javax", "activation", xs @ _*) => MergeStrategy.last
  case PathList("org", "apache", xs @ _*) => MergeStrategy.last
  case PathList("com", "google", xs @ _*) => MergeStrategy.last
  case "plugin.properties" => MergeStrategy.last
  case "log4j.properties" => MergeStrategy.last
  case x =>
    val oldStrategy = (assembly / assemblyMergeStrategy).value
    oldStrategy(x)
}

Compile / PB.targets := Seq(
  scalapb.gen() -> (Compile / sourceManaged).value / "scalapb"
)
Compile / PB.protoSources += file("../protobuf")

assemblyJarName := { s"${name.value}-${version.value}.jar" }
