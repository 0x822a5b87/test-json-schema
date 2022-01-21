package v1

type StreamSource struct {
	Type        StreamSourceTypeEnum `json:"type" validate:"required"`
	KafkaSpecs  KafkaSpec            `json:"kafkaSpec,omitempty"`
	PulsarSpecs PulsarSpec           `json:"pulsarSpec,omitempty"`
}

type KafkaSpec struct {
	// kafkaID
	Id        string            `json:"id" validate:"required"`
	Topic     string            `json:"topic" validate:"required"`
	Zookeeper string            `json:"zookeeper" validate:"required"`
	Broker    string            `json:"broker" validate:"required"`
	Config    map[string]string `json:"config"`
}

type PulsarSpec struct {
}
