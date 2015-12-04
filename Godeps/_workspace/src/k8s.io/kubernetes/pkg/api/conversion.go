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

package api

import (
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/conversion"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/runtime"
)

// Codec is the identity codec for this package - it can only convert itself
// to itself.
var Codec = runtime.CodecFor(Scheme, "")

func init() {
	Scheme.AddDefaultingFuncs(
		func(obj *unversioned.ListOptions) {
			if obj.LabelSelector.Selector == nil {
				obj.LabelSelector = unversioned.LabelSelector{labels.Everything()}
			}
			if obj.FieldSelector.Selector == nil {
				obj.FieldSelector = unversioned.FieldSelector{fields.Everything()}
			}
		},
	)
	Scheme.AddConversionFuncs(
		func(in *unversioned.Time, out *unversioned.Time, s conversion.Scope) error {
			// Cannot deep copy these, because time.Time has unexported fields.
			*out = *in
			return nil
		},
		func(in *string, out *labels.Selector, s conversion.Scope) error {
			selector, err := labels.Parse(*in)
			if err != nil {
				return err
			}
			*out = selector
			return nil
		},
		func(in *string, out *fields.Selector, s conversion.Scope) error {
			selector, err := fields.ParseSelector(*in)
			if err != nil {
				return err
			}
			*out = selector
			return nil
		},
		func(in *labels.Selector, out *string, s conversion.Scope) error {
			if *in == nil {
				return nil
			}
			*out = (*in).String()
			return nil
		},
		func(in *fields.Selector, out *string, s conversion.Scope) error {
			if *in == nil {
				return nil
			}
			*out = (*in).String()
			return nil
		},
		func(in *string, out *unversioned.LabelSelector, s conversion.Scope) error {
			selector, err := labels.Parse(*in)
			if err != nil {
				return err
			}
			*out = unversioned.LabelSelector{selector}
			return nil
		},
		func(in *string, out *unversioned.FieldSelector, s conversion.Scope) error {
			selector, err := fields.ParseSelector(*in)
			if err != nil {
				return err
			}
			*out = unversioned.FieldSelector{selector}
			return nil
		},
		func(in *[]string, out *unversioned.LabelSelector, s conversion.Scope) error {
			selectorString := ""
			if len(*in) > 0 {
				selectorString = (*in)[0]
			}
			selector, err := labels.Parse(selectorString)
			if err != nil {
				return err
			}
			*out = unversioned.LabelSelector{selector}
			return nil
		},
		func(in *[]string, out *unversioned.FieldSelector, s conversion.Scope) error {
			selectorString := ""
			if len(*in) > 0 {
				selectorString = (*in)[0]
			}
			selector, err := fields.ParseSelector(selectorString)
			if err != nil {
				return err
			}
			*out = unversioned.FieldSelector{selector}
			return nil
		},
		func(in *unversioned.LabelSelector, out *string, s conversion.Scope) error {
			if in.Selector == nil {
				return nil
			}
			*out = in.Selector.String()
			return nil
		},
		func(in *unversioned.FieldSelector, out *string, s conversion.Scope) error {
			if in.Selector == nil {
				return nil
			}
			*out = in.Selector.String()
			return nil
		},
		func(in *unversioned.LabelSelector, out *unversioned.LabelSelector, s conversion.Scope) error {
			if in.Selector == nil {
				return nil
			}
			selector, err := labels.Parse(in.Selector.String())
			if err != nil {
				return err
			}
			out.Selector = selector
			return nil
		},
		func(in *unversioned.FieldSelector, out *unversioned.FieldSelector, s conversion.Scope) error {
			if in.Selector == nil {
				return nil
			}
			selector, err := fields.ParseSelector(in.Selector.String())
			if err != nil {
				return err
			}
			out.Selector = selector
			return nil
		},
		func(in *resource.Quantity, out *resource.Quantity, s conversion.Scope) error {
			// Cannot deep copy these, because inf.Dec has unexported fields.
			*out = *in.Copy()
			return nil
		},
	)
}
