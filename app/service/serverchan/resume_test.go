package serverchan

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestResumeRemind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ResumeRemind()
	})
}
