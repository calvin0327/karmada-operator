// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/daocloud/karmada-operator/pkg/generated/clientset/versioned/typed/install/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeInstallV1alpha1 struct {
	*testing.Fake
}

func (c *FakeInstallV1alpha1) KarmadaDeployments() v1alpha1.KarmadaDeploymentInterface {
	return &FakeKarmadaDeployments{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeInstallV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
