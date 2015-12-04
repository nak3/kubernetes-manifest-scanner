/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package validation

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/util/validation"
)

// ValidateEvent makes sure that the event makes sense.
func ValidateEvent(event *api.Event) validation.ErrorList {
	allErrs := validation.ErrorList{}
	// There is no namespace required for node.
	if event.InvolvedObject.Kind == "Node" &&
		event.Namespace != "" {
		allErrs = append(allErrs, validation.NewInvalidError(validation.NewFieldPath("involvedObject", "namespace"), event.InvolvedObject.Namespace, "namespace is not required for node"))
	}
	if event.InvolvedObject.Kind != "Node" &&
		event.Namespace != event.InvolvedObject.Namespace {
		allErrs = append(allErrs, validation.NewInvalidError(validation.NewFieldPath("involvedObject", "namespace"), event.InvolvedObject.Namespace, "does not match involvedObject"))
	}
	if !validation.IsDNS1123Subdomain(event.Namespace) {
		allErrs = append(allErrs, validation.NewInvalidError(validation.NewFieldPath("namespace"), event.Namespace, ""))
	}
	return allErrs
}
