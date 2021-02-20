package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
