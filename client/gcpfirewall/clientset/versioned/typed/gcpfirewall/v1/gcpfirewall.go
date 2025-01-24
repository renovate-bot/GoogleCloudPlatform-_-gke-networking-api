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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/gcpfirewall/v1"
	scheme "github.com/GoogleCloudPlatform/gke-networking-api/client/gcpfirewall/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// GCPFirewallsGetter has a method to return a GCPFirewallInterface.
// A group's client should implement this interface.
type GCPFirewallsGetter interface {
	GCPFirewalls() GCPFirewallInterface
}

// GCPFirewallInterface has methods to work with GCPFirewall resources.
type GCPFirewallInterface interface {
	Create(ctx context.Context, gCPFirewall *v1.GCPFirewall, opts metav1.CreateOptions) (*v1.GCPFirewall, error)
	Update(ctx context.Context, gCPFirewall *v1.GCPFirewall, opts metav1.UpdateOptions) (*v1.GCPFirewall, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, gCPFirewall *v1.GCPFirewall, opts metav1.UpdateOptions) (*v1.GCPFirewall, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GCPFirewall, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GCPFirewallList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GCPFirewall, err error)
	GCPFirewallExpansion
}

// gCPFirewalls implements GCPFirewallInterface
type gCPFirewalls struct {
	*gentype.ClientWithList[*v1.GCPFirewall, *v1.GCPFirewallList]
}

// newGCPFirewalls returns a GCPFirewalls
func newGCPFirewalls(c *NetworkingV1Client) *gCPFirewalls {
	return &gCPFirewalls{
		gentype.NewClientWithList[*v1.GCPFirewall, *v1.GCPFirewallList](
			"gcpfirewalls",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.GCPFirewall { return &v1.GCPFirewall{} },
			func() *v1.GCPFirewallList { return &v1.GCPFirewallList{} }),
	}
}
