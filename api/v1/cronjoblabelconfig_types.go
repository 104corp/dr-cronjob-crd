/*
Copyright 2024 deep huang.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronJobLabelConfigSpec defines the desired state of CronJobLabelConfig.
type CronJobLabelConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CronJobLabelConfig. Edit cronjoblabelconfig_types.go to remove/update
	EnableCleanup bool   `json:"enableCleanup,omitempty"`
	Foo           string `json:"foo,omitempty"`
}

// CronJobLabelConfigStatus defines the observed state of CronJobLabelConfig.
type CronJobLabelConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	SyncedCronJobs int `json:"syncedCronJobs"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CronJobLabelConfig is the Schema for the cronjoblabelconfigs API.
type CronJobLabelConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobLabelConfigSpec   `json:"spec,omitempty"`
	Status CronJobLabelConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronJobLabelConfigList contains a list of CronJobLabelConfig.
type CronJobLabelConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJobLabelConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJobLabelConfig{}, &CronJobLabelConfigList{})
}
