/*
Copyright 2023.

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

// ResponderSpec defines the desired state of Responder
type ResponderSpec struct{}

// ResponderStatus defines the observed state of Responder
type ResponderStatus struct {
	Acquired bool `json:"acquired"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Responder is the Schema for the responders API
type Responder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResponderSpec   `json:"spec,omitempty"`
	Status ResponderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ResponderList contains a list of Responder
type ResponderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Responder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Responder{}, &ResponderList{})
}
