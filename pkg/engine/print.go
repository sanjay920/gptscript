package engine

import (
	"strings"

	"github.com/sanjay920/gptscript/pkg/counter"
	"github.com/sanjay920/gptscript/pkg/types"
)

func (e *Engine) runEcho(tool types.Tool) (cmdOut *Return, cmdErr error) {
	id := counter.Next()
	out := strings.TrimPrefix(tool.Instructions, types.EchoPrefix+"\n")

	e.Progress <- types.CompletionStatus{
		CompletionID: id,
		Response: map[string]any{
			"output": out,
			"err":    nil,
		},
	}

	return &Return{
		Result: &out,
	}, nil
}
