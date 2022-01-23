package v1

type Sink struct {
	Type  SinkTypeEnum `yaml:"type" validate:"required"`
	Redis RedisSink    `yaml:"redis"`
}

type RedisSink struct {
	Host string `yaml:"host" validate:"required"`
	Port uint   `yaml:"port" validate:"required"`
	Db   string `yaml:"db" validate:"required"`
	Auth string `yaml:"auth" validate:"required"`
	// Redis 使用资源的匹配模式
	Pattern KeyPattern `yaml:"pattern" validate:"required"`

	Config map[string]string `yaml:"config"`
}

type KeyPattern struct {
	Prefixes []string `yaml:"prefixes"`
}
