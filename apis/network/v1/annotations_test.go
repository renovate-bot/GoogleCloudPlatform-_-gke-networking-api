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
package v1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/utils/ptr"
)

func TestPodIPsAnnotation(t *testing.T) {
	tests := []struct {
		name     string
		input    PodIPsAnnotation
		expected string
	}{
		{
			name:     "nil",
			input:    nil,
			expected: "null",
		},
		{
			name:     "empty list",
			input:    PodIPsAnnotation{},
			expected: "[]",
		},
		{
			name: "single pod IP",
			input: PodIPsAnnotation{
				{NetworkName: "network-a", IP: "198.51.100.0"},
			},
			expected: `[{"networkName":"network-a","ip":"198.51.100.0"}]`,
		},
		{
			name: "missing network",
			input: PodIPsAnnotation{
				{IP: "198.51.100.0"},
			},
			expected: `[{"networkName":"","ip":"198.51.100.0"}]`,
		},
		{
			name: "multiple pod IPs",
			input: PodIPsAnnotation{
				{NetworkName: "network-a", IP: "198.51.100.0"},
				{NetworkName: "network-b", IP: "2001:db8::"},
			},
			expected: `[{"networkName":"network-a","ip":"198.51.100.0"},{"networkName":"network-b","ip":"2001:db8::"}]`,
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

			parsed, err := ParsePodIPsAnnotation(marshalled)
			if err != nil {
				t.Fatalf("ParsePodIPsAnnotation(%s) failed with error: %v", marshalled, err)
			}

			if diff := cmp.Diff(parsed, tc.input); diff != "" {
				t.Fatalf("ParsePodIPsAnnotation(%s) returns diff: (-got +want): %s", marshalled, diff)
			}
		})
	}
}

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
		{
			name: "list with items with subnets",
			input: NodeNetworkAnnotation{
				{Name: "network-a", IPv4Subnet: "198.51.100.0/24", IPv6Subnet: "2001:db8::/32"},
				{Name: "network-b", IPv4Subnet: "198.52.100.0/24", IPv6Subnet: "2001:db9::/32"},
			},
			expected: `[{"name":"network-a","ipv4-subnet":"198.51.100.0/24","ipv6-subnet":"2001:db8::/32"},{"name":"network-b","ipv4-subnet":"198.52.100.0/24","ipv6-subnet":"2001:db9::/32"}]`,
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

func TestParseMultiNetworkAnnotation(t *testing.T) {
	tests := []struct {
		name     string
		input    MultiNetworkAnnotation
		expected string
	}{
		{
			name:     "nil",
			input:    nil,
			expected: "null",
		},
		{
			name:     "empty list",
			input:    MultiNetworkAnnotation{},
			expected: "[]",
		},
		{
			name: "list with items",
			input: MultiNetworkAnnotation{
				{Name: "network-a", Cidrs: []string{"1.1.1.1/21"}, Scope: "host-local"},
				{Name: "network-b", Cidrs: []string{"2.2.2.2/12"}, Scope: "global"},
			},
			expected: `[{"name":"network-a","cidrs":["1.1.1.1/21"],"scope":"host-local"},{"name":"network-b","cidrs":["2.2.2.2/12"],"scope":"global"}]`,
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

			parsed, err := ParseMultiNetworkAnnotation(marshalled)
			if err != nil {
				t.Fatalf("ParseMultiNetworkAnnotation(%s) failed with error: %v", marshalled, err)
			}

			if diff := cmp.Diff(parsed, tc.input); diff != "" {
				t.Fatalf("ParseMultiNetworkAnnotation(%s) returns diff: (-got +want): %s", marshalled, diff)
			}
		})
	}
}

func TestParseNorthInterfacesAnnotation(t *testing.T) {
	tests := []struct {
		name     string
		input    NorthInterfacesAnnotation
		expected string
	}{
		{
			name:     "nil",
			input:    nil,
			expected: "null",
		},
		{
			name:     "empty list",
			input:    NorthInterfacesAnnotation{},
			expected: "[]",
		},
		{
			name: "list with items",
			input: NorthInterfacesAnnotation{
				{Network: "network-a", IpAddress: "10.0.0.1"},
				{Network: "network-b", IpAddress: "20.0.0.1"},
			},
			expected: `[{"network":"network-a","ipAddress":"10.0.0.1"},{"network":"network-b","ipAddress":"20.0.0.1"}]`,
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

			parsed, err := ParseNorthInterfacesAnnotation(marshalled)
			if err != nil {
				t.Fatalf("ParseNorthInterfacesAnnotation(%s) failed with error: %v", marshalled, err)
			}

			if diff := cmp.Diff(parsed, tc.input); diff != "" {
				t.Fatalf("ParseNorthInterfacesAnnotation(%s) returns diff: (-got +want): %s", marshalled, diff)
			}
		})
	}
}

func TestParseNICInfoAnnotation(t *testing.T) {
	tests := []struct {
		name     string
		input    NICInfoAnnotation
		expected string
	}{
		{
			name:     "nil",
			input:    nil,
			expected: "null",
		},
		{
			name:     "empty list",
			input:    NICInfoAnnotation{},
			expected: "[]",
		},
		{
			name: "list with items",
			input: NICInfoAnnotation{
				{BirthIP: "10.0.0.1", PCIAddress: "0000:00:05.0", BirthName: "eth1"},
				{BirthIP: "20.0.0.1", PCIAddress: "0000:00:06.0", BirthName: "eth2"},
			},
			expected: `[{"birthIP":"10.0.0.1","pciAddress":"0000:00:05.0","birthName":"eth1"},{"birthIP":"20.0.0.1","pciAddress":"0000:00:06.0","birthName":"eth2"}]`,
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

			parsed, err := ParseNICInfoAnnotation(marshalled)
			if err != nil {
				t.Fatalf("ParseNICInfoAnnotation(%s) failed with error: %v", marshalled, err)
			}

			if diff := cmp.Diff(parsed, tc.input); diff != "" {
				t.Fatalf("ParseNICInfoAnnotation(%s)  returns diff: (-got +want): %s", marshalled, diff)
			}
		})
	}
}

func TestInterfaceStatusAnnotation(t *testing.T) {
	tests := []struct {
		name     string
		input    InterfaceStatusAnnotation
		expected string
	}{
		{
			name:     "nil",
			input:    nil,
			expected: "null",
		},
		{
			name:     "empty list",
			input:    InterfaceStatusAnnotation{},
			expected: "[]",
		},
		{
			name: "minimal status",
			input: InterfaceStatusAnnotation{
				InterfaceStatus{
					NetworkName: "network",
					IPAddresses: []string{"1.2.3.4"},
					MACAddress:  "aa:bb:cc:dd:ee",
				},
			},
			expected: `[{"networkName":"network","ipAddresses":["1.2.3.4"],"macAddress":"aa:bb:cc:dd:ee"}]`,
		},
		{
			name: "max status",
			input: InterfaceStatusAnnotation{
				InterfaceStatus{
					NetworkName: "network",
					IPAddresses: []string{"1.2.3.4", "ff::01"},
					MACAddress:  "aa:bb:cc:dd:ee",
					Routes: []Route{
						{To: "10.0.0.0/24"},
					},
					Gateway4: ptr.To("10.0.0.1"),
					DNSConfig: &DNSConfig{
						Nameservers: []string{"8.8.8.8"},
						Searches:    []string{"a.domain"},
					},
					DHCPServerIP: ptr.To("10.0.0.2"),
				},
			},
			expected: `[{"networkName":"network","ipAddresses":["1.2.3.4","ff::01"],"macAddress":"aa:bb:cc:dd:ee","routes":[{"to":"10.0.0.0/24"}],"gateway4":"10.0.0.1","dnsConfig":{"nameservers":["8.8.8.8"],"searches":["a.domain"]},"dhcpServerIP":"10.0.0.2"}]`,
		},
		{
			name: "multiple interfaces",
			input: InterfaceStatusAnnotation{
				InterfaceStatus{
					NetworkName: "network",
					IPAddresses: []string{"1.2.3.4"},
					MACAddress:  "aa:bb:cc:dd:ee",
				},
				InterfaceStatus{
					NetworkName: "network-2",
					IPAddresses: []string{"1.2.3.5", "ff::01"},
					MACAddress:  "aa:bb:cc:dd:ff",
					Routes: []Route{
						{To: "10.0.0.0/24"},
					},
					Gateway4: ptr.To("10.0.0.1"),
					DNSConfig: &DNSConfig{
						Nameservers: []string{"8.8.8.8"},
						Searches:    []string{"a.domain"},
					},
					DHCPServerIP: ptr.To("10.0.0.2"),
				},
			},
			expected: `[{"networkName":"network","ipAddresses":["1.2.3.4"],"macAddress":"aa:bb:cc:dd:ee"},{"networkName":"network-2","ipAddresses":["1.2.3.5","ff::01"],"macAddress":"aa:bb:cc:dd:ff","routes":[{"to":"10.0.0.0/24"}],"gateway4":"10.0.0.1","dnsConfig":{"nameservers":["8.8.8.8"],"searches":["a.domain"]},"dhcpServerIP":"10.0.0.2"}]`,
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

			parsed, err := ParseInterfaceStatusAnnotation(marshalled)
			if err != nil {
				t.Fatalf("ParseInterfaceStatusAnnotation(%s) failed with error: %v", marshalled, err)
			}

			if diff := cmp.Diff(parsed, tc.input); diff != "" {
				t.Fatalf("ParseInterfaceStatusAnnotation(%s) returns diff: (-got +want): %s", marshalled, diff)
			}
		})
	}
}
