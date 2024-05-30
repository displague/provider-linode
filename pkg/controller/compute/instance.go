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

package compute

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/displague/crossplane-provider-linode/apis/compute/v1alpha1"
	apisv1alpha1 "github.com/displague/crossplane-provider-linode/apis/v1alpha1"
)

const (
	errNotInstance                = "managed resource is not a Instance custom resource"
	errProviderNotRetrieved       = "provider could not be retrieved"
	errProviderSecretNil          = "cannot find Secret reference on Provider"
	errProviderSecretNotRetrieved = "secret referred in provider could not be retrieved"

	errNewClient = "cannot create new Service"
)

// A NoOpService does nothing.
type NoOpService struct{}

var (
	newNoOpService = func() (interface{}, error) { return &NoOpService{}, nil }
)

// SetupInstance adds a controller that reconciles Instance managed resources.
func SetupInstance(mgr ctrl.Manager, l logging.Logger) error {
	name := managed.ControllerName(v1alpha1.InstanceGroupKind)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.InstanceGroupVersionKind),
		managed.WithExternalConnecter(&instanceConnector{kube: mgr.GetClient(), newServiceFn: newNoOpService}),
		managed.WithInitializers(managed.NewNameAsExternalName(mgr.GetClient())),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&v1alpha1.Instance{}).
		Complete(r)
}

type instanceConnector struct {
	kube         client.Client
	newServiceFn func() (interface{}, error)
}

func (c *instanceConnector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.Instance)
	if !ok {
		return nil, errors.New(errNotInstance)
	}

	provider := &apisv1alpha1.Provider{}
	if err := c.kube.Get(ctx, meta.NamespacedNameOf(cr.Spec.ProviderReference), provider); err != nil {
		return nil, errors.Wrap(err, errProviderNotRetrieved)
	}

	if provider.GetCredentialsSecretReference() == nil {
		return nil, errors.New(errProviderSecretNil)
	}

	secret := &v1.Secret{}
	n := types.NamespacedName{Namespace: provider.Spec.CredentialsSecretRef.Namespace, Name: provider.Spec.CredentialsSecretRef.Name}
	if err := c.kube.Get(ctx, n, secret); err != nil {
		return nil, errors.Wrap(err, errProviderSecretNotRetrieved)
	}

	s, err := c.newServiceFn()
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	return &instanceExternal{kube: c.kube, service: s}, nil
}

type instanceExternal struct {
	kube    client.Client
	service interface{}
}

func (c *instanceExternal) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.Instance)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotInstance)
	}

	fmt.Printf("Observing: %+v", cr)

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
		// ConnectionDetails: getConnectionDetails(cr, instance),
	}, nil
}

func (c *instanceExternal) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.Instance)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotInstance)
	}

	fmt.Printf("Creating: %+v", cr)

	return managed.ExternalCreation{}, nil
}

func (c *instanceExternal) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.Instance)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotInstance)
	}

	fmt.Printf("Updating: %+v", cr)

	return managed.ExternalUpdate{}, nil
}

func (c *instanceExternal) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.Instance)
	if !ok {
		return errors.New(errNotInstance)
	}

	fmt.Printf("Deleting: %+v", cr)

	return nil
}

func createRandomRootPassword() (string, error) {
	rawRootPass := make([]byte, 50)
	_, err := rand.Read(rawRootPass)
	if err != nil {
		return "", fmt.Errorf("Failed to generate random password")
	}
	rootPass := base64.StdEncoding.EncodeToString(rawRootPass)
	return rootPass, nil
}
