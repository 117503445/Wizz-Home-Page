package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
