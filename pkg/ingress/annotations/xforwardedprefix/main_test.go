/*
Copyright 2017 The Kubernetes Authors.

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

package xforwardedprefix

import (
	"testing"

	"github.com/stolostron/management-ingress/pkg/ingress/annotations/parser"
	"github.com/stolostron/management-ingress/pkg/ingress/resolver"
	api "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestParse(t *testing.T) {
	annotation := parser.GetAnnotationWithPrefix("x-forwarded-prefix")
	ap := NewParser(&resolver.Mock{})
	if ap == nil {
		t.Fatalf("expected a parser.IngressAnnotation but returned nil")
	}

	testCases := []struct {
		annotations map[string]string
		expected    bool
	}{
		{map[string]string{annotation: "true"}, true},
		{map[string]string{annotation: "1"}, true},
		{map[string]string{annotation: ""}, false},
		{map[string]string{}, false},
		{nil, false},
	}

	ing := &networking.Ingress{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
		},
		Spec: networking.IngressSpec{},
	}

	for _, testCase := range testCases {
		ing.SetAnnotations(testCase.annotations)
		result, _ := ap.Parse(ing)
		if result != testCase.expected {
			t.Errorf("expected %v but returned %v, annotations: %s", testCase.expected, result, testCase.annotations)
		}
	}
}
