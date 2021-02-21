package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestGetCurrent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		interview := GetCurrentInterview()
		g.Log().Line().Debug(interview)
	})
}
