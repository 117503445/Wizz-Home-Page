package wjx

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestGetSojumpSurvey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		GetSojumpSurvey()
		if SojumpSurvey, err := gcache.Get("SojumpSurvey"); err != nil {
			g.Log().Line().Error(err)
		} else {
			t.Assert(len(SojumpSurvey.(string)) > 0, true)
		}
	})
}
