/*
Copyright 2016 The Kubernetes Authors.

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

package secureupstream

import (
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"

	"k8s.io/ingress/core/pkg/ingress/annotations/parser"
)

const (
	secureUpstream = "ingress.kubernetes.io/secure-backends"
)

type su struct {
}

// NewParser creates a new secure upstream annotation parser
func NewParser() parser.IngressAnnotation {
	return su{}
}

// Parse parses the annotations contained in the ingress
// rule used to indicate if the upstream servers should use SSL
func (a su) Parse(ing *extensions.Ingress) (interface{}, error) {
	return parser.GetBoolAnnotation(secureUpstream, ing)
}
