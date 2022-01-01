ThisBuild / version := "1.0"
ThisBuild / scalaVersion := "2.12.10"
ThisBuild / organization := "net.jp.ytake"

name := "split-brain-groups"
logLevel := Level.Error
resolvers ++= Seq(
  "mvn" at "https://mvnrepository.com/artifact"
)

libraryDependencies ++= Seq(
  "org.scala-lang" % "scala-library" % "2.12.10",
  "com.typesafe" % "config" % "1.2.0",
  "org.apache.spark" %% "spark-core" % "3.0.3" % "provided",
  "org.apache.spark" %% "spark-sql" % "3.0.3" % "provided",
  "org.apache.spark" %% "spark-sql-kafka-0-10" % "3.0.3",
  "com.thesamet.scalapb" %% "sparksql-scalapb" % "0.11.0",
  "com.thesamet.scalapb" %% "scalapb-runtime" % scalapb.compiler.Version.scalapbVersion % "protobuf",
  "junit" % "junit" % "4.13" % Test,
  "org.scalatest" %% "scalatest" % "3.2.7" % Test
)

assembly / assemblyShadeRules := Seq(
  ShadeRule.rename("com.google.protobuf.**" -> "shadeproto.@1").inAll,
  ShadeRule.rename("scala.collection.compat.**" -> "scalacompat.@1").inAll
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
