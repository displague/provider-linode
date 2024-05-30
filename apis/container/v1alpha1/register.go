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
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "container.linode.crossplane.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// LKECluster type metadata.
var (
	LKEClusterKind             = reflect.TypeOf(LKECluster{}).Name()
	LKEClusterGroupKind        = schema.GroupKind{Group: Group, Kind: LKEClusterKind}.String()
	LKEClusterKindAPIVersion   = LKEClusterKind + "." + SchemeGroupVersion.String()
	LKEClusterGroupVersionKind = SchemeGroupVersion.WithKind(LKEClusterKind)
)

// LKEClusterClass type metadata.
var (
	LKEClusterClassKind             = reflect.TypeOf(LKEClusterClass{}).Name()
	LKEClusterClassGroupKind        = schema.GroupKind{Group: Group, Kind: LKEClusterClassKind}.String()
	LKEClusterClassKindAPIVersion   = LKEClusterClassKind + "." + SchemeGroupVersion.String()
	LKEClusterClassGroupVersionKind = SchemeGroupVersion.WithKind(LKEClusterClassKind)
)

// LKEClusterPool type metadata.
var (
	LKEClusterPoolKind             = reflect.TypeOf(LKEClusterPool{}).Name()
	LKEClusterPoolGroupKind        = schema.GroupKind{Group: Group, Kind: LKEClusterPoolKind}.String()
	LKEClusterPoolKindAPIVersion   = LKEClusterPoolKind + "." + SchemeGroupVersion.String()
	LKEClusterPoolGroupVersionKind = SchemeGroupVersion.WithKind(LKEClusterPoolKind)
)

// LKEClusterPoolClass type metadata.
var (
	LKEClusterPoolClassKind             = reflect.TypeOf(LKEClusterPoolClass{}).Name()
	LKEClusterPoolClassGroupKind        = schema.GroupKind{Group: Group, Kind: LKEClusterPoolClassKind}.String()
	LKEClusterPoolClassKindAPIVersion   = LKEClusterPoolClassKind + "." + SchemeGroupVersion.String()
	LKEClusterPoolClassGroupVersionKind = SchemeGroupVersion.WithKind(LKEClusterPoolClassKind)
)

func init() {
	SchemeBuilder.Register(&LKECluster{}, &LKEClusterList{})
	SchemeBuilder.Register(&LKEClusterClass{}, &LKEClusterClassList{})
	SchemeBuilder.Register(&LKEClusterPool{}, &LKEClusterPoolList{})
	SchemeBuilder.Register(&LKEClusterPoolClass{}, &LKEClusterPoolClassList{})
}
