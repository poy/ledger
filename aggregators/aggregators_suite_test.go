package aggregators_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAggregators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Aggregators Suite")
}
