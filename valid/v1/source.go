package v1

type StreamSource struct {
	Type       StreamSourceTypeEnum `yaml:"type" validate:"required"`
	KafkaSpec  KafkaSpec            `yaml:"kafka" validate:"required"`
	PulsarSpec PulsarSpec           `yaml:"pulsar"`
}

type KafkaSpec struct {
	// kafkaID
	Id        string            `yaml:"id" validate:"required"`
	Topic     string            `yaml:"topic" validate:"required"`
	Zookeeper string            `yaml:"zookeeper" validate:"required"`
	Broker    string            `yaml:"broker" validate:"required"`
	Config    map[string]string `yaml:"config"`
}

type PulsarSpec struct {
}
