package wjx

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestDownload(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		Download()
		ParseExcel()
	})
}
