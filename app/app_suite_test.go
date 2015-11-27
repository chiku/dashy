package app_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}

type BadReadCloser struct{ err error }

func (rc *BadReadCloser) Close() error             { return nil }
func (rc *BadReadCloser) Read([]byte) (int, error) { return 0, rc.err }
