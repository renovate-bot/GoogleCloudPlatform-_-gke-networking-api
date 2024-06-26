/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNodeNetworkAnnotation(t *testing.T) {
	tests := []struct {
		name     string
		input    NodeNetworkAnnotation
		expected string
	}{
		{
			name:     "nil",
			input:    nil,
			expected: "null",
		},
		{
			name:     "empty list",
			input:    NodeNetworkAnnotation{},
			expected: "[]",
		},
		{
			name: "list with items",
			input: NodeNetworkAnnotation{
				{Name: "network-a"},
				{Name: "network-b"},
			},
			expected: `[{"name":"network-a"},{"name":"network-b"}]`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			marshalled, err := MarshalAnnotation(tc.input)
			if err != nil {
				t.Fatalf("MarshalAnnotation(%+v) failed with error: %v", tc.input, err)
			}
			if marshalled != tc.expected {
				t.Fatalf("MarshalAnnotation(%+v) returns %q but want %q", tc.input, marshalled, tc.expected)
			}

			parsed, err := ParseNodeNetworkAnnotation(marshalled)
			if err != nil {
				t.Fatalf("ParseNodeNetworkAnnotation(%s) failed with error: %v", marshalled, err)
			}

			if diff := cmp.Diff(parsed, tc.input); diff != "" {
				t.Fatalf("ParseNodeNetworkAnnotation(%s) returns diff: (-got +want): %s", marshalled, diff)
			}
		})
	}
}
