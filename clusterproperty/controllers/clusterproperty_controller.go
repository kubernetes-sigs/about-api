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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	aboutv1alpha1 "k8s.io/clusterproperty/api/v1alpha1"
)

// ClusterPropertyReconciler reconciles a ClusterProperty object
type ClusterPropertyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=about.k8s.io,resources=clusterproperties,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=about.k8s.io,resources=clusterproperties/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=about.k8s.io,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=about.k8s.io,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=about.k8s.io,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=api.autoscaling.v2beta1,resources=hpa,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=about.k8s.io,resources=clusterproperties/finalizers,verbs=update

func (r *ClusterPropertyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var clusterProperty aboutv1alpha1.ClusterProperty

	// TODO: which CRs in the CRD is watching
	// Reconcilitation watch the CR with name specifically "id.k8s.io"

	if err := r.Get(ctx, req.NamespacedName, &clusterProperty); err != nil {
		log.Log.Error(err, "Unable to fetch the required clusterProperty")
		// we will ignore the not found errors, since they cannot be fixed by an immediate
		// requeue( we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// if err := r.Status().Update(ctx, &clusterProperty); err != nil {
	// 	log.Log.Error(err, "unable to update ClusterProperty status")
	// 	return ctrl.Result{}, err
	// }

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterPropertyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&aboutv1alpha1.ClusterProperty{}).
		Complete(r)
}
