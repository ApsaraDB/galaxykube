/*
Copyright 2021 Alibaba Group Holding Limited.

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

package control

import (
	"time"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func wait(name, msg string) BindFunc {
	return NewStepBinder(NewStep(name,
		func(rc ReconcileContext, flow Flow) (reconcile.Result, error) {
			return flow.Wait(msg)
		}),
	)
}

func Wait(msg string) BindFunc {
	return wait("Wait", msg)
}

func Abort(msg string) BindFunc {
	return wait("Abort", msg)
}

func AbortWhen(cond bool, msg string) BindFunc {
	return When(cond, Abort(msg))
}

// Requeue equivalent to Retry.
var Requeue = Retry

func Retry(msg string) BindFunc {
	return NewStepBinder(NewStep("Retry",
		func(rc ReconcileContext, flow Flow) (reconcile.Result, error) {
			return flow.Retry(msg)
		}),
	)
}

// RequeueAfter equivalent to RetryAfter
var RequeueAfter = RetryAfter

func RetryAfter(d time.Duration, msg string) BindFunc {
	return NewStepBinder(NewStep("RetryAfter"+d.String(),
		func(rc ReconcileContext, flow Flow) (reconcile.Result, error) {
			return flow.RetryAfter(d, msg)
		}),
	)
}

func ScheduleAfter(d time.Duration) BindFunc {
	return NewStepBinder(NewStep("ScheduleAfter"+d.String(),
		func(rc ReconcileContext, flow Flow) (reconcile.Result, error) {
			rc.ResetForceRequeueAfter(d)
			return flow.Pass()
		}),
	)
}
