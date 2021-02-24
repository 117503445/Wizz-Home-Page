package serverchan

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestAlarm(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		Alarm("测试ALARM:服务器出问题啦", "测试Content")
	})
}
