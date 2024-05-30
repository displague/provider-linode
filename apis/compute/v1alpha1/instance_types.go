/*
Copyright 2019 The Crossplane Authors.

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

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// Defaults for Instance resources.
const (
	DefaultReclaimPolicy = runtimev1alpha1.ReclaimRetain
)

// +kubebuilder:validation:Required

// InstanceStatus defines the observed state of Instance
type InstanceObservation struct {
	// ID is the unique immutable numeric identifier of a Linode Instance
	// +optional
	ID int `json:"id,omitempty"`

	// Status is the current activity status of a Linode Instance
	Status string `json:"status"`

	// Label is the unique mutable name of a Linode Instance
	Label string `json:"label"`

	// Region defines the geographic location of a Linode Instance
	Region string `json:"region"`

	// Type is the Linode Instance Type which represents the cost, processor, memory, transfer, and storage profile of the Instance
	Type string `json:"type"`

	// IPv6 is the public IPv6 address of a Linode Instance
	// +optional
	IPv6 string `json:"ipv6,omitempty"`

	// IPv4 is the list of IPv4 addresses associated with a Linode Instance
	// +optional
	IPv4 []string `json:"ipv4,omitempty"`

	// Image is the image detected on a Linode Instance disk
	// +optional
	Image string `json:"image,omitempty"`
}

type InstanceParameters struct {
	// Label is the unique name of this Linode Instance
	// +optional
	Label string `json:"label,omitempty"`

	// Image is the disk image to be applied to the first instance disk
	// +optional
	Image string `json:"image,omitempty"`

	// AuthorizedUsers are Linode user accounts whose SSH keys will be authorized to SSH into the instance
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty"`

	// Region defines the geographic location of a Linode Instance
	Region string `json:"region"`

	// Type is the Linode Instance Type which represents the cost, processor, memory, transfer, and storage profile of the Instance
	Type string `json:"type"`

	// Status is the current activity status of a Linode Instance
	// +kubebuilder:validation:Enum=offline;running
	Status string `json:"status,omitempty"`
}

// A InstanceSpec defines the desired state of a Instance.
type InstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  InstanceParameters `json:"forProvider"`
}

// A InstanceStatus represents the observed state of a Instance.
type InstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Instance is a managed resource that represents a Linode Kubernetes Engine
// node pool.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="CLUSTER-NAME",type="string",JSONPath=".spec.forProvider.cluster"
// +kubebuilder:printcolumn:name="NODE-POOL-CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".spec.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instance items
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// A InstanceClassSpecTemplate is a template for the spec of a dynamically
// provisioned Instance.
type InstanceClassSpecTemplate struct {
	runtimev1alpha1.ClassSpecTemplate `json:",inline"`
	InstanceParameters                `json:",inline"`
}

// +kubebuilder:object:root=true

// A InstanceClass is a resource class. It defines the desired spec of
// resource claims that use it to dynamically provision a managed
// resource.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="PROVIDER-REF",type="string",JSONPath=".specTemplate.providerRef.name"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".specTemplate.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,class,linode}
type InstanceClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// SpecTemplate is a template for the spec of a dynamically provisioned
	// Instance.
	SpecTemplate InstanceClassSpecTemplate `json:"specTemplate"`
}

// +kubebuilder:object:root=true

// InstanceClassList contains a list of cloud memorystore resource classes.
type InstanceClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceClass `json:"items"`
}
