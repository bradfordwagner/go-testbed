package internal

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrivateMock", func() {
	It("should do stuff", func() { Expect(true).To(BeTrue()) })
	It("unmocked returns hi friends", func() {
		p := new(PrivateMock)
		Expect(p.SayHello()).To(Equal("hi friends"))
	})
})
