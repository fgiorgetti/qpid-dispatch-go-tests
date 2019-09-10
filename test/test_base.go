package test

import (
	"fmt"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework/ginkgowrapper"
	"github.com/interconnectedcloud/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// Unique and synchronized Setup
var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Unique initialization (node 1 only)
	fmt.Println("Base test setup - only happens once per test suite")
	return nil
}, func(data []byte) {
	// Initilization for each parallel node
}, 10)


// Unique and synchronized Teardown
var _ = ginkgo.SynchronizedAfterSuite(func() {
	// All nodes tear down
}, func() {
	// Node1 only tear down
	fmt.Println("Base test teardown - only happens once per test suite")
}, 10)

// CreateInterconnect creates an Interconnect resource using the provided InterconnectSpec
func CreateInterconnect(c *framework.ContextData, size int32, spec v1alpha1.InterconnectSpec) (*v1alpha1.Interconnect, error) {
	return c.CreateInterconnect(c.Namespace, size, func(interconnect *v1alpha1.Interconnect) {
		interconnect.Spec = spec
	})
}

func init() {
	framework.HandleFlags()
	gomega.RegisterFailHandler(ginkgowrapper.Fail)
}

func ContextsAvailable() int {
	return len(framework.TestContext.GetContexts())
}
