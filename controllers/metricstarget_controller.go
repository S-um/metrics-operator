/*
Copyright 2021.

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

	klog "log"

	zerooneaiv1 "github.com/myeongsuk.yoon/metrics-operator/api/v1"
	"github.com/myeongsuk.yoon/metrics-operator/prometheus/metricsRouter"
)

// MetricsTargetReconciler reconciles a MetricsTarget object
type MetricsTargetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

var ReconcileLoopCnt int = 0

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MetricsTarget object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
//+kubebuilder:rbac:groups=zeroone.ai.myeongsukyoon,resources=metricstargets,verbs=get;list;watch;create;update;patch;delete;watch
//+kubebuilder:rbac:groups=zeroone.ai.myeongsukyoon,resources=metricstargets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=zeroone.ai.myeongsukyoon,resources=metricstargets/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;
func (r *MetricsTargetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	targetConfig := &zerooneaiv1.MetricsTarget{}
	r.Get(ctx, req.NamespacedName, targetConfig)

	klog.Println("Updating Custom Resource")
	klog.Println("targetConfig.Name :", targetConfig.Name)
	metricsRouter.UpdateConfig(*targetConfig, req.Name)
	klog.Println("Success to Update Custom Resource")
	klog.Println("Reconcile :", ReconcileLoopCnt)
	klog.Println(req)
	klog.Println(ctx)
	ReconcileLoopCnt++
	klog.Println(targetConfig)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MetricsTargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&zerooneaiv1.MetricsTarget{}).
		Complete(r)
}
