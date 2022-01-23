package v1

type JobTypeEnum string

const (
	Flink JobTypeEnum = "flink"
	Storm JobTypeEnum = "storm"
)

type StreamSourceTypeEnum string

const (
	SourceKafka  StreamSourceTypeEnum = "kafka"
	SourcePulsar StreamSourceTypeEnum = "pulsar"
)

type DeployTypeEnum string

const (
	Jar       DeployTypeEnum = "jar"
	Container DeployTypeEnum = "container"
)

type SinkTypeEnum string

const (
	Redis SinkTypeEnum = "redis"
	Kafka SinkTypeEnum = "kafka"
)

// TypeMeta job metadata
type TypeMeta struct {
	Kind       JobTypeEnum `yaml:"kind" validate:"required"`
	APIVersion string      `yaml:"apiVersion" validate:"required"`
}

type ObjectMeta struct {
	Name string `yaml:"name" validate:"required"`

	Labels map[string]string `yaml:"labels"`

	GenerateName string `yaml:"generateName,omitempty"`

	Namespace string `yaml:"namespace,omitempty"`
}

type Deploy struct {
	DeployType DeployTypeEnum `yaml:"type" validate:"required"`

	JarDeploy `yaml:"jar" validate:"required"`
}

type JarDeploy struct {
	JarName string `yaml:"name" validate:"required"`
	Class   string `yaml:"class" validate:"required"`
}
