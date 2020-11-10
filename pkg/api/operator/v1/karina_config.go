package v1

import (
	"github.com/flanksource/karina/pkg/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KarinaConfigSpec defines the desired state of KarinaConfig
type KarinaConfigSpec struct {
	DryRun       bool                      `json:"dryRun,omitempty"`
	Config       types.PlatformConfig      `json:"config,omitempty"`
	TemplateFrom map[string]TemplateSource `json:"templateFrom,omitempty"`
}

// KarinaConfigStatus defines the observed state of KarinaConfig
type KarinaConfigStatus struct {
	LastApplied         metav1.Time `json:"lastApplied,omitempty"`
	LastAppliedChecksum string      `json:"lastAppliedChecksum,omitempty"`
}

type TemplateSource struct {
	// Write the content of secret/configmap/template to a file
	// and set field to file name
	// +optional
	Tmpfile bool `json:"tmpFile,omitempty"`
	// Applies a Golang template
	// +optional
	Template *TemplateSourceValue `json:"templateValue,omitempty"`
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *v1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	// Selects a key of a secret in the pod's namespace
	// +optional
	SecretKeyRef *v1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

type TemplateSourceValue struct {
	Template string `json:"template,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KarinaConfig is the Schema for the KarinaConfigs API
type KarinaConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KarinaConfigSpec   `json:"spec,omitempty"`
	Status KarinaConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KarinaConfigList contains a list of KarinaConfig
type KarinaConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KarinaConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KarinaConfig{}, &KarinaConfigList{})
}
