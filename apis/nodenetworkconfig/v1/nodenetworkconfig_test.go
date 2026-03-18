/*
Copyright 2026 The Kubernetes Authors.

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

package v1

import (
	"encoding/json"
	"testing"

	"k8s.io/apimachinery/pkg/util/strategicpatch"
)

func TestNodeNetworkConfig(t *testing.T) {
	nnc := &NodeNetworkConfig{
		Spec: NodeNetworkConfigSpec{
			Allocations: []Allocation{
				{
					Network: DefaultPodNetworkName,
					Pods:    100,
				},
			},
		},
	}

	if len(nnc.Spec.Allocations) != 1 {
		t.Fatalf("expected 1 allocation, got %d", len(nnc.Spec.Allocations))
	}

	allocation := nnc.Spec.Allocations[0]
	if allocation.Network != DefaultPodNetworkName {
		t.Errorf("expected network %q, got %q", DefaultPodNetworkName, allocation.Network)
	}
	if allocation.Pods != 100 {
		t.Errorf("expected pods 100, got %d", allocation.Pods)
	}
}

func TestNodeNetworkConfigReleasableCIDRsPatch(t *testing.T) {
	original := &NodeNetworkConfig{
		Spec: NodeNetworkConfigSpec{
			ReleasableCIDRs: []PodCIDR{
				{Id: "id-1", Network: "net-1", CIDR: "10.0.0.0/24"},
			},
		},
	}

	patchData := []byte(`{"spec":{"releasableCIDRs":[{"id":"id-2","network":"net-2","cidr":"10.0.1.0/24"},{"id":"id-1","network":"net-1-patched","cidr":"10.0.0.0/24"}]}}`)

	originalJSON, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("failed to marshal original: %v", err)
	}

	patchedJSON, err := strategicpatch.StrategicMergePatch(originalJSON, patchData, &NodeNetworkConfig{})
	if err != nil {
		t.Fatalf("failed to apply strategic merge patch: %v", err)
	}

	var patched NodeNetworkConfig
	if err := json.Unmarshal(patchedJSON, &patched); err != nil {
		t.Fatalf("failed to unmarshal patched object: %v", err)
	}

	if len(patched.Spec.ReleasableCIDRs) != 2 {
		t.Errorf("expected 2 ReleasableCIDRs after patch, got %d", len(patched.Spec.ReleasableCIDRs))
	}

	for _, c := range patched.Spec.ReleasableCIDRs {
		switch c.Id {
		case "id-1":
			if c.Network != "net-1-patched" {
				t.Errorf("expected id-1 network to be merged to 'net-1-patched', got %q", c.Network)
			}
		case "id-2":
			if c.Network != "net-2" {
				t.Errorf("expected id-2 network to be appended with 'net-2', got %q", c.Network)
			}
		default:
			t.Errorf("unexpected podCIDR id: %q", c.Id)
		}
	}
}
