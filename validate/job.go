package validate

import "json/validate/v1"

type StormJob struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty"`

	StormSpec `json:"spec" validate:"required"`
}

type FlinkJob struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty"`

	FlinkSpec `json:"spec" validate:"required"`
}

type StormSpec struct {
	Sources []v1.StreamSource `json:"sources" validate:"required"`
	Sink    v1.Sink           `json:"sink" validate:"required"`

	// Probes are not allowed for ephemeral containers.
	// +optional
	LivenessProbe *Probe `json:"livenessProbe,omitempty"`
	// Probes are not allowed for ephemeral containers.
	// +optional
	ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
}

type FlinkSpec struct {
	// EnableCheckpoint
	EnableCheckpoint bool `json:"enableCheckpoint,omitempty"`

	Labels map[string]string `json:"labels"`

	// Probes are not allowed for ephemeral containers.
	// +optional
	LivenessProbe *Probe `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	// Probes are not allowed for ephemeral containers.
	// +optional
	ReadinessProbe *Probe `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
}

// Probe describes a health check to be performed against a container to determine whether it is
// alive or ready to receive traffic.
type Probe struct {
}
