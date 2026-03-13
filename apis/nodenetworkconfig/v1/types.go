/*
Copyright 2026 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

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

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=nnc,scope=Cluster
// +kubebuilder:storageversion
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkConfig describes the network configuration for a Node.
// +k8s:openapi-gen=true
type NodeNetworkConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeNetworkConfigSpec   `json:"spec,omitempty"`
	Status NodeNetworkConfigStatus `json:"status,omitempty"`
}

// NodeNetworkConfigSpec is the spec for a NodeNetworkConfig resource.
// +k8s:openapi-gen=true
type NodeNetworkConfigSpec struct {
	// Allocations is a list of network allocations.
	// +optional
	// +listType=atomic
	Allocations []Allocation `json:"allocations,omitempty"`

	// ReleasableCIDRs is a list of releasable pod CIDRs.
	// +optional
	// +listType=atomic
	ReleasableCIDRs []PodCIDR `json:"releasableCIDRs,omitempty"`
}

// Allocation describes the network allocation for a specific network.
// +k8s:openapi-gen=true
type Allocation struct {
	// Network is the name of the network.
	// +required
	Network string `json:"network"`

	// Pods is the number of pods allocated from this network.
	// +optional
	Pods int32 `json:"pods,omitempty"`
}

// PodCIDR describes a pod CIDR.
// +k8s:openapi-gen=true
type PodCIDR struct {
	// Id is the identifier of the pod CIDR.
	// +required
	Id string `json:"id"`

	// Network is the name of the network.
	// +required
	Network string `json:"network"`

	// CIDR is the pod CIDR range.
	// +required
	CIDR string `json:"cidr"`

	// Conditions contains details for the current condition of this pod CIDR.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// NodeNetworkConfigStatus is the status for a NodeNetworkConfig resource.
// +k8s:openapi-gen=true
type NodeNetworkConfigStatus struct {
	// PodCIDRs is a list of pod CIDRs.
	// +optional
	// +listType=atomic
	PodCIDRs []PodCIDR `json:"podCIDRs,omitempty"`

	// Conditions contains details for the current condition of this resource.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkConfigList contains a list of NodeNetworkConfig resources.
type NodeNetworkConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is a list of NodeNetworkConfig.
	Items []NodeNetworkConfig `json:"items"`
}
