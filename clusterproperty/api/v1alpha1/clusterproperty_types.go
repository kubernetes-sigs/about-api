/*
Copyright 2022 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterPropertySpec defines the desired state of ClusterProperty
type ClusterPropertySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ClusterProperty Name
	Name string `json:"name,omitempty"`

	// ClusterProperty value
	// +kubebuilder:validation:Maxlength=128000
	// +kubebuilder:validation:MinLength=1
	Value string `json:"value"`
}

// ClusterPropertyStatus defines the observed state of ClusterProperty
type ClusterPropertyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// ClusterProperty is the Schema for the clusterproperties API
// +kubebuilder:printcolumn:name="value",type=string,JSONPath=`.spec.value`
// +kubebuilder:printcolumn:name="age",type=date,JSONPath=`.metadata.creationTimestamp`
type ClusterProperty struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterPropertySpec   `json:"spec,omitempty"`
	Status ClusterPropertyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterPropertyList contains a list of ClusterProperty
type ClusterPropertyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterProperty `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterProperty{}, &ClusterPropertyList{})
}
