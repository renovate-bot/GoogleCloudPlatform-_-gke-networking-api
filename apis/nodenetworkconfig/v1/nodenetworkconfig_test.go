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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
)

func TestNodeNetworkConfig(t *testing.T) {
	nnc := &NodeNetworkConfig{
		Spec: NodeNetworkConfigSpec{
			Allocations: []Allocation{
				{
					Pods: 100,
				},
			},
		},
		Status: NodeNetworkConfigStatus{
			PodCIDRs: []PodCIDR{
				{
					Id:      "id-test",
					Network: "default",
					CIDR:    "10.0.0.0/24",
					Condition: &metav1.Condition{
						Type:   string(PodCIDRConditionReady),
						Status: metav1.ConditionTrue,
					},
				},
			},
			Conditions: []metav1.Condition{
				{
					Type:   string(NodeNetworkConfigConditionReady),
					Status: metav1.ConditionTrue,
					Reason: string(NodeNetworkConfigInvalidParametersReason),
				},
			},
		},
	}

	if len(nnc.Spec.Allocations) != 1 {
		t.Fatalf("expected 1 allocation, got %d", len(nnc.Spec.Allocations))
	}

	allocation := nnc.Spec.Allocations[0]
	// Explicitly test for empty string because the standard Go behavior defaults missing string fields to "",
	// whereas the +kubebuilder:default annotation only takes effect when the object traverses the Kubernetes API Server.
	if allocation.Network != "" {
		t.Errorf("expected network %q, got %q", "", allocation.Network)
	}
	if allocation.Pods != 100 {
		t.Errorf("expected pods 100, got %d", allocation.Pods)
	}

	if len(nnc.Status.PodCIDRs) != 1 {
		t.Fatalf("expected 1 podCIDR, got %d", len(nnc.Status.PodCIDRs))
	}
	podCIDR := nnc.Status.PodCIDRs[0]
	if podCIDR.Id != "id-test" {
		t.Errorf("expected podCIDR id 'id-test', got %q", podCIDR.Id)
	}
	if podCIDR.Network != "default" {
		t.Errorf("expected podCIDR network %q, got %q", "default", podCIDR.Network)
	}
	if podCIDR.CIDR != "10.0.0.0/24" {
		t.Errorf("expected podCIDR CIDR '10.0.0.0/24', got %q", podCIDR.CIDR)
	}
	if podCIDR.Condition == nil || podCIDR.Condition.Type != string(PodCIDRConditionReady) || podCIDR.Condition.Status != metav1.ConditionTrue {
		t.Errorf("expected podCIDR condition Ready=True, got %v", podCIDR.Condition)
	}

	if len(nnc.Status.Conditions) != 1 {
		t.Fatalf("expected 1 condition, got %d", len(nnc.Status.Conditions))
	}
	condition := nnc.Status.Conditions[0]
	if condition.Type != string(NodeNetworkConfigConditionReady) {
		t.Errorf("expected condition type %q, got %q", NodeNetworkConfigConditionReady, condition.Type)
	}
	if condition.Status != metav1.ConditionTrue {
		t.Errorf("expected condition status True, got %q", condition.Status)
	}
	if condition.Reason != string(NodeNetworkConfigInvalidParametersReason) {
		t.Errorf("expected condition reason %q, got %q", NodeNetworkConfigInvalidParametersReason, condition.Reason)
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
