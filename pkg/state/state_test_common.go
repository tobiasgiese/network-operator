/*
2022 NVIDIA CORPORATION & AFFILIATES

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

package state

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type checkFunc func(object *unstructured.Unstructured)

func runFuncForObjectInSlice(objects []*unstructured.Unstructured, objectKind string, f checkFunc) bool {
	var found bool
	for _, obj := range objects {
		if obj.GetKind() != objectKind {
			continue
		}
		found = true

		f(obj)
	}
	return found
}
