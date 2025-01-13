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

package fake

import (
	v1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/nodetopology/v1"
	nodetopologyv1 "github.com/GoogleCloudPlatform/gke-networking-api/client/nodetopology/clientset/versioned/typed/nodetopology/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeNodeTopologies implements NodeTopologyInterface
type fakeNodeTopologies struct {
	*gentype.FakeClientWithList[*v1.NodeTopology, *v1.NodeTopologyList]
	Fake *FakeNetworkingV1
}

func newFakeNodeTopologies(fake *FakeNetworkingV1) nodetopologyv1.NodeTopologyInterface {
	return &fakeNodeTopologies{
		gentype.NewFakeClientWithList[*v1.NodeTopology, *v1.NodeTopologyList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("nodetopologies"),
			v1.SchemeGroupVersion.WithKind("NodeTopology"),
			func() *v1.NodeTopology { return &v1.NodeTopology{} },
			func() *v1.NodeTopologyList { return &v1.NodeTopologyList{} },
			func(dst, src *v1.NodeTopologyList) { dst.ListMeta = src.ListMeta },
			func(list *v1.NodeTopologyList) []*v1.NodeTopology { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.NodeTopologyList, items []*v1.NodeTopology) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
