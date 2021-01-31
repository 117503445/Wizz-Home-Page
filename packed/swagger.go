package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
