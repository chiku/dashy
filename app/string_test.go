package app_test

import (
	a "github.com/chiku/dashy/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String position in slice", func() {
	Context("when searched in slice containing the item", func() {
		It("is the first occurance position", func() {
			Expect(a.StringPosInSlice("abc", []string{"a", "b", "abc"})).To(Equal(2))
		})
	})

	Context("when searched in slice not containing the item", func() {
		It("is -1", func() {
			Expect(a.StringPosInSlice("abc", []string{"a", "b", "ABC"})).To(Equal(-1))
		})
	})

	Context("when searched in an empty slice", func() {
		It("is -1", func() {
			Expect(a.StringPosInSlice("abc", []string{})).To(Equal(-1))
		})
	})
})
