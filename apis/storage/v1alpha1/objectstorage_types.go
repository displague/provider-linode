/*
Copyright 2020 The Crossplane Authors.

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
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Required

// ObjectStorageParameters defines the desired state of Lke
type ObjectStorageParameters struct {
	Label   string `json:"label"`
	Region  string `json:"region"`
	Version string `json:"version"`

	// +optional
	Tags []string `json:"tags,omitempty"`
}

// ObjectStorageObservation is used to show the observed state of the LKE Cluster
// resource on Linode.
type ObjectStorageObservation struct {
	ID      int         `json:"id"`
	Created metav1.Time `json:"created"`
	Updated metav1.Time `json:"updated"`
	Label   string      `json:"label"`
	Region  string      `json:"region"`
	Status  string      `json:"status"`
	Version string      `json:"version"`
	Tags    []string    `json:"tags"`
}

// A ObjectStorageSpec defines the desired state of a ObjectStorage.
type ObjectStorageSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ObjectStorageParameters `json:"forProvider"`
}

// A ObjectStorageStatus represents the observed state of a ObjectStorage.
type ObjectStorageStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ObjectStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A ObjectStorage is a managed resource that represents a Linode Kubernetes Engine
// node pool.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="CLUSTER-NAME",type="string",JSONPath=".spec.forProvider.cluster"
// +kubebuilder:printcolumn:name="NODE-POOL-CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".spec.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type ObjectStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectStorageSpec   `json:"spec"`
	Status ObjectStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectStorageList contains a list of ObjectStorage items
type ObjectStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectStorage `json:"items"`
}

// A ObjectStorageClassSpecTemplate is a template for the spec of a dynamically
// provisioned ObjectStorage.
type ObjectStorageClassSpecTemplate struct {
	runtimev1alpha1.ClassSpecTemplate `json:",inline"`
	ObjectStorageParameters           `json:",inline"`
}

// +kubebuilder:object:root=true

// A ObjectStorageClass is a resource class. It defines the desired spec of
// resource claims that use it to dynamically provision a managed
// resource.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="PROVIDER-REF",type="string",JSONPath=".specTemplate.providerRef.name"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".specTemplate.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,class,linode}
type ObjectStorageClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// SpecTemplate is a template for the spec of a dynamically provisioned
	// ObjectStorage.
	SpecTemplate ObjectStorageClassSpecTemplate `json:"specTemplate"`
}

// +kubebuilder:object:root=true

// ObjectStorageClassList contains a list of cloud memorystore resource classes.
type ObjectStorageClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectStorageClass `json:"items"`
}
