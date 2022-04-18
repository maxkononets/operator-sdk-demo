/*
Copyright 2022.

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

// DemoControllerSpec defines the desired state of DemoController
type DemoControllerSpec struct {
	Size int32 `json:"size"`
}

// DemoControllerStatus defines the observed state of DemoController
type DemoControllerStatus struct {
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DemoController is the Schema for the democontrollers API
//+kubebuilder:subresource:status
type DemoController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoControllerSpec   `json:"spec,omitempty"`
	Status DemoControllerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DemoControllerList contains a list of DemoController
type DemoControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoController `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoController{}, &DemoControllerList{})
}
