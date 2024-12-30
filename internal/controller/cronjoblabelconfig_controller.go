/*
Copyright 2024 deep huang.

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

package controller

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"

	appv1 "104corp.org/dr-cronjob-crd/api/v1"
	batchv1 "k8s.io/api/batch/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CronJobLabelConfigReconciler reconciles a CronJobLabelConfig object
type CronJobLabelConfigReconciler struct {
	client.Client
}

// +kubebuilder:rbac:groups=batch.104corp.org,resources=cronjoblabelconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.104corp.org,resources=cronjoblabelconfigs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch.104corp.org,resources=cronjoblabelconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CronJobLabelConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *CronJobLabelConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var config appv1.CronJobLabelConfig
	if err := r.Get(ctx, req.NamespacedName, &config); err != nil {
		if errors.IsNotFound(err) {
			// 資源被刪除，不需要處理
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// 獲取所有 CronJob
	var cronJobs batchv1.CronJobList
	if err := r.List(ctx, &cronJobs, &client.ListOptions{Namespace: req.Namespace}); err != nil {
		return ctrl.Result{}, err
	}

	// 更新 CronJob 標籤
	syncedCount := 0
	for _, cronJob := range cronJobs.Items {
		labels := cronJob.GetLabels()
		if labels == nil {
			labels = make(map[string]string)
		}

		// 根據 EnableCleanup 決定是否加上/移除標籤
		if config.Spec.EnableCleanup {
			if labels["cleanup"] != "true" {
				labels["cleanup"] = "true"
				cronJob.SetLabels(labels)
				if err := r.Update(ctx, &cronJob); err != nil {
					return ctrl.Result{}, err
				}
			}
		} else {
			if _, exists := labels["cleanup"]; exists {
				delete(labels, "cleanup")
				cronJob.SetLabels(labels)
				if err := r.Update(ctx, &cronJob); err != nil {
					return ctrl.Result{}, err
				}
			}
		}

		syncedCount++
	}

	// 更新 Status
	config.Status.SyncedCronJobs = syncedCount
	if err := r.Status().Update(ctx, &config); err != nil {
		return ctrl.Result{}, err
	}

	fmt.Printf("Synced %d CronJobs with cleanup label\n", syncedCount)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CronJobLabelConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.CronJobLabelConfig{}).
		Named("cronjoblabelconfig").
		Complete(r)
}
