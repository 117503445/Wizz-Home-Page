package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
