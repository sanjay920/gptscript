package embedded

import (
	"io/fs"
	"os"

	"github.com/sanjay920/gptscript/internal"
	"github.com/sanjay920/gptscript/pkg/cli"
	"github.com/sanjay920/gptscript/pkg/system"
)

type Options struct {
	FS fs.FS
}

func Run(opts ...Options) bool {
	for _, opt := range opts {
		if opt.FS != nil {
			internal.FS = opt.FS
		}
	}

	system.SetBinToSelf()
	if os.Getenv("GPTSCRIPT_EMBEDDED") == "true" {
		cli.Main()
		return true
	}

	_ = os.Setenv("GPTSCRIPT_EMBEDDED", "true")
	return false
}
