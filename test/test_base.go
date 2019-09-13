package test

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework/ginkgowrapper"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// Initialize Ginkgo and parse command line arguments
func Initialize() {
	framework.HandleFlags()
	gomega.RegisterFailHandler(ginkgowrapper.Fail)
}

// Before suite validation setup (happens only once per test suite)
var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Unique initialization (node 1 only)
	return nil
}, func(data []byte) {
	// Initialization for each parallel node
}, 10)

// After suite validation teardown (happens only once per test suite)
var _ = ginkgo.SynchronizedAfterSuite(func() {
	// All nodes tear down
}, func() {
	// Node1 only tear down
	framework.RunCleanupActions(framework.AfterEach)
	framework.RunCleanupActions(framework.AfterSuite)
}, 10)
