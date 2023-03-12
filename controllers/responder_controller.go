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

package controllers

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/controller-e2e-framework/test-2-controller/api/v1alpha1"
	deliveryv1alpha1 "github.com/controller-e2e-framework/test-2-controller/api/v1alpha1"
)

// ResponderReconciler reconciles a Responder object
type ResponderReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=delivery.controller-e2e-framework,resources=responders,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=delivery.controller-e2e-framework,resources=responders/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=delivery.controller-e2e-framework,resources=responders/finalizers,verbs=update
//+kubebuilder:rbac:groups=delivery.controller-e2e-framework,resources=controller,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Responder object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ResponderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	obj := &v1alpha1.Responder{}

	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("failed to get responder: %w", err)
	}

	if obj.Status.Controlled {
		logger.Info("this object has been acquired by a controller")
		// Initialize the patch helper.
		patchHelper, err := patch.NewHelper(obj, r.Client)
		if err != nil {
			return ctrl.Result{}, err
		}
		obj.Status.Acquired = true
		if err := patchHelper.Patch(ctx, obj); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ResponderReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&deliveryv1alpha1.Responder{}).
		Complete(r)
}
