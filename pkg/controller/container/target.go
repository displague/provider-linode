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

package container

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/target"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	workloadv1alpha1 "github.com/crossplane/crossplane/apis/workload/v1alpha1"

	"github.com/displague/crossplane-provider-linode/apis/container/v1alpha1"
)

// SetupLKEClusterTarget adds a controller that propagates LKECluster
// connection secrets to the connection secrets of their targets.
func SetupLKEClusterTarget(mgr ctrl.Manager, l logging.Logger) error {
	name := target.ControllerName(v1alpha1.LKEClusterGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&workloadv1alpha1.KubernetesTarget{}).
		WithEventFilter(resource.NewPredicates(resource.HasManagedResourceReferenceKind(resource.ManagedKind(v1alpha1.LKEClusterGroupVersionKind)))).
		Complete(target.NewReconciler(mgr,
			resource.TargetKind(workloadv1alpha1.KubernetesTargetGroupVersionKind),
			resource.ManagedKind(v1alpha1.LKEClusterGroupVersionKind),
			target.WithLogger(l.WithValues("controller", name)),
			target.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}
