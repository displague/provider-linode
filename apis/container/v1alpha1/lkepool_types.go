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

	"github.com/linode/linodego"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LKEClusterPool states.
const (
	LKEClusterPoolStateUnspecified  = "STATUS_UNSPECIFIED"
	LKEClusterPoolStateProvisioning = "PROVISIONING"
	LKEClusterPoolStateRunning      = "RUNNING"
	LKEClusterPoolStateRunningError = "RUNNING_WITH_ERROR"
	LKEClusterPoolStateReconciling  = "RECONCILING"
	LKEClusterPoolStateStopping     = "STOPPING"
	LKEClusterPoolStateError        = "ERROR"
)

// Error strings
const (
	errResourceIsNotLKEClusterPool = "the managed resource is not a LKEClusterPool"
)

// LKEClusterPoolObservation is used to show the observed state of the LKE Node Pool
// resource on Linode.
type LKEClusterPoolObservation struct {
	// The Node Pool's unique ID.
	ID int `json:"id,omitempty"`

	// The Linode Type for all nodes in the Node Pool
	Type string `json:"type,omitempty"`

	// The number of nodes in the Node Pool.
	Count int `json:"count,omitempty"`

	// Status information for the Nodes which are members of this Node Pool. If a Linode has not been provisioned for a given Node slot, the instance_id will be returned as null.
	Nodes []linodego.LKEClusterPoolLinode `json:"nodes,omitempty"`
}

// LKEClusterPoolParameters define the desired state of a Linode Kubernetes
// Engine node pool.
type LKEClusterPoolParameters struct {
	// NOTE(displague): Cluster is marked as omitempty but is not optional. It
	// will either be assigned a value directly or set from the ClusterRef.

	// Cluster: The LKE Cluster ID for the LKE cluster to which the
	// LKEClusterPool will attach. Must be of format Must be supplied if
	// ClusterRef is not.
	// +immutable
	Cluster string `json:"cluster,omitempty"`

	// ClusterRef sets the Cluster field by resolving the resource link of the
	// referenced Crossplane LKECluster managed resource. Must be supplied in
	// Cluster is not.
	// +immutable
	// +optional
	ClusterRef *runtimev1alpha1.Reference `json:"clusterRef,omitempty"`

	// ClusterSelector selects a reference to resolve the resource link of the
	// referenced Crossplane LKECluster managed resource.
	// +immutable
	// +optional
	ClusterSelector *runtimev1alpha1.Selector `json:"clusterSelector,omitempty"`

	// The number of nodes in the Node Pool.
	Count int `json:"count,omitempty"`

	// A Linode Type for all of the nodes in the Node Pool.
	// +immutable
	Type string `json:"type,omitempty"`
}

// A LKEClusterPoolSpec defines the desired state of a LKEClusterPool.
type LKEClusterPoolSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LKEClusterPoolParameters `json:"forProvider"`
}

// A LKEClusterPoolStatus represents the observed state of a LKEClusterPool.
type LKEClusterPoolStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LKEClusterPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A LKEClusterPool is a managed resource that represents a Linode Kubernetes Engine
// node pool.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="CLUSTER-NAME",type="string",JSONPath=".spec.forProvider.cluster"
// +kubebuilder:printcolumn:name="NODE-POOL-CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".spec.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type LKEClusterPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LKEClusterPoolSpec   `json:"spec"`
	Status LKEClusterPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LKEClusterPoolList contains a list of LKEClusterPool items
type LKEClusterPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LKEClusterPool `json:"items"`
}

// A LKEClusterPoolClassSpecTemplate is a template for the spec of a dynamically
// provisioned LKEClusterPool.
type LKEClusterPoolClassSpecTemplate struct {
	runtimev1alpha1.ClassSpecTemplate `json:",inline"`
	LKEClusterPoolParameters          `json:",inline"`
}

// +kubebuilder:object:root=true

// A LKEClusterPoolClass is a resource class. It defines the desired spec of
// resource claims that use it to dynamically provision a managed
// resource.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="PROVIDER-REF",type="string",JSONPath=".specTemplate.providerRef.name"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".specTemplate.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,class,linode}
type LKEClusterPoolClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// SpecTemplate is a template for the spec of a dynamically provisioned
	// LKEClusterPool.
	SpecTemplate LKEClusterPoolClassSpecTemplate `json:"specTemplate"`
}

// +kubebuilder:object:root=true

// LKEClusterPoolClassList contains a list of cloud memorystore resource classes.
type LKEClusterPoolClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LKEClusterPoolClass `json:"items"`
}
