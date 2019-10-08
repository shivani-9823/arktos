/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAuthorizationV1 struct {
	*testing.Fake
}

func (c *FakeAuthorizationV1) LocalSubjectAccessReviews(namespace string, optional_tenant ...string) v1.LocalSubjectAccessReviewInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeLocalSubjectAccessReviews{c, namespace, tenant}
}

func (c *FakeAuthorizationV1) SelfSubjectAccessReviews(optional_tenant ...string) v1.SelfSubjectAccessReviewInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeSelfSubjectAccessReviews{c, tenant}
}

func (c *FakeAuthorizationV1) SelfSubjectRulesReviews(optional_tenant ...string) v1.SelfSubjectRulesReviewInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeSelfSubjectRulesReviews{c, tenant}
}

func (c *FakeAuthorizationV1) SubjectAccessReviews(optional_tenant ...string) v1.SubjectAccessReviewInterface {
	tenant := "default"
	if len(optional_tenant) > 0 {
		tenant = optional_tenant[0]
	}
	return &FakeSubjectAccessReviews{c, tenant}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAuthorizationV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
