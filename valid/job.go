package valid

import "json/valid/v1"

type StormJob struct {
	v1.TypeMeta   `yaml:",inline"`
	v1.ObjectMeta `yaml:"metadata,omitempty"`

	StormSpec `yaml:"spec" validate:"required"`
}

type FlinkJob struct {
	v1.TypeMeta   `yaml:",inline"`
	v1.ObjectMeta `yaml:"metadata,omitempty"`

	FlinkSpec `yaml:"spec" validate:"required"`
}

type StormSpec struct {
	Sources []v1.StreamSource `yaml:"sources" validate:"required"`
	Sink    v1.Sink           `yaml:"sink" validate:"required"`

	// Probes are not allowed for ephemeral containers.
	// +optional
	LivenessProbe *Probe `yaml:"livenessProbe,omitempty"`
	// Probes are not allowed for ephemeral containers.
	// +optional
	ReadinessProbe *Probe `yaml:"readinessProbe,omitempty"`
}

type FlinkSpec struct {
	v1.Deploy `yaml:"deploy" validate:"required"`

	Sources []v1.StreamSource `yaml:"sources" validate:"required"`
	Sink    v1.Sink           `yaml:"sink" validate:"required"`

	Config FlinkJobConfig `yaml:"config"`

	Labels map[string]string `yaml:"labels"`

	// Probes are not allowed for ephemeral containers.
	// +optional
	LivenessProbe *Probe `yaml:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	// Probes are not allowed for ephemeral containers.
	// +optional
	ReadinessProbe *Probe `yaml:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
}

// Probe describes a health check to be performed against a container to determine whether it is
// alive or ready to receive traffic.
type Probe struct {
}

type FlinkJobConfig struct {
	// EnableCheckpoint
	EnableCheckpoint bool `yaml:"enableCheckpoint"`

	// TODO flink 资源，并行度等配置
}
