package v1

type RedisKeyPatternEnum string

const (
	Regex  RedisKeyPatternEnum = "regex"
	Prefix RedisKeyPatternEnum = "prefix"
)

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
	Type RedisKeyPatternEnum `yaml:"type" validate:"required"`

	Prefixes []string `yaml:"prefixes"`
	Regex    []string `yaml:"regexes"`
}

func NewKeyPattern() KeyPattern {
	return KeyPattern{
		Type: Prefix,
	}
}
