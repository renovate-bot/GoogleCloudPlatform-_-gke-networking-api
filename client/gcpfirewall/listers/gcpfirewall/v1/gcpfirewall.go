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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/gcpfirewall/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// GCPFirewallLister helps list GCPFirewalls.
// All objects returned here must be treated as read-only.
type GCPFirewallLister interface {
	// List lists all GCPFirewalls in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GCPFirewall, err error)
	// Get retrieves the GCPFirewall from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GCPFirewall, error)
	GCPFirewallListerExpansion
}

// gCPFirewallLister implements the GCPFirewallLister interface.
type gCPFirewallLister struct {
	listers.ResourceIndexer[*v1.GCPFirewall]
}

// NewGCPFirewallLister returns a new GCPFirewallLister.
func NewGCPFirewallLister(indexer cache.Indexer) GCPFirewallLister {
	return &gCPFirewallLister{listers.New[*v1.GCPFirewall](indexer, v1.Resource("gcpfirewall"))}
}
