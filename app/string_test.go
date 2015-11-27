package app_test

import (
	a "github.com/chiku/dashy/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String", func() {
	Context("When searched in slice containing the item", func() {
		It("Exists", func() {
			Expect(a.StringInSlice("abc", []string{"a", "b", "abc"})).To(BeTrue())
		})
	})

	Context("When searched in slice not containing the item", func() {
		It("Doesn't exist", func() {
			Expect(a.StringInSlice("abc", []string{"a", "b", "ABC"})).To(BeFalse())
		})
	})

	Context("When searched in an empty slice", func() {
		It("Doesn't exist", func() {
			Expect(a.StringInSlice("abc", []string{})).To(BeFalse())
		})
	})
})
