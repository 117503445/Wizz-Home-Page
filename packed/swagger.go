package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
