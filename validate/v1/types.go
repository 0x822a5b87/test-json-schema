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

// TypeMeta job metadata
type TypeMeta struct {
	Kind       JobTypeEnum `json:"kind" validate:"required"`
	APIVersion string      `json:"apiVersion" validate:"required"`
}

type ObjectMeta struct {
	Name string `json:"name" validate:"required"`

	Labels map[string]string `json:"labels"`

	GenerateName string `json:"generateName,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type Deploy struct {
	DeployType DeployTypeEnum `json:"type" validate:"required"`

	JarDeploy
}

type JarDeploy struct {
	JarName string `json:"jar" validate:"required"`
	Class   string `json:"class" validate:"required"`
}
