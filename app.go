package main

import (
	"json/validate"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
kind: flink
# 版本号
apiVersion: v1
metadata:
  name: flink-kafka-write2-es-xxx
  namespace: default
  # 标签
  labels:
    # 任务需要在上海跑
    location: sh
    # 表示任务需要DB写入权限
    authority: db_write
# 任务描述
spec:
  # flink 任务的专属选项，是否开启
  enable-checkpoint: true
  type: jar
  # jar包
  jar: FlinkStream.jar
  # 类名
  class: com.tencent.stream.KafkaWrite2RedisJob
  source:
    - type: kafka
      id: 202
      topic: ieg_smoba_playerlogin
      zookeeper: 127.0.0.1:2181
      broker: 127.0.0.1:9092
      config:
        security.protocol: org.apache.kafka.common.security.scram.ScramLoginModule
        sasl.mechanism: sasl-username
        sasl.jaas.password: sasl-password
    - type: kafka
      id: 203
      topic: ieg_smoba_playerlogin
    - type: kafka
      id: 204
      topic: ieg_smoba_playerlogin
    - type: kafka
      id: 205
      topic: ieg_smoba_playerlogin
  sink:
    - type: redis
      host: tencent-redis
      port: 6379
      db: smoba
      auth: xxx
  # 监控配置
  monitor:
    - type: tglog
      host: tglog.sh
      port: 33088
  # 状态检查
  probe:
`

func main() {
	flinkJob := validate.FlinkJob{}
	err := yaml.Unmarshal([]byte(data), &flinkJob)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}