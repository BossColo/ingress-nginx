/*
Copyright 2015 The Kubernetes Authors.

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

package store

import (
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// ServiceLister makes a Store that lists Services.
type IngressResourceLister struct {
	cache.Store
}

// ByKey returns the Service matching key in the local Service Store.
func (sl *IngressResourceLister) ByKey(key string) (*extensionsv1beta1.Ingress, error) {
	i, exists, err := sl.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, NotExistsError(key)
	}
	return i.(*extensionsv1beta1.Ingress), nil
}
