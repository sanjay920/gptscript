package builtin

import (
	"github.com/sanjay920/gptscript/pkg/openai"
	"github.com/sanjay920/gptscript/pkg/types"
)

var (
	defaultModel = openai.DefaultModel
)

func GetDefaultModel() string {
	return defaultModel
}

func SetDefaultModel(model string) {
	defaultModel = model
}

func SetDefaults(tool types.Tool) types.Tool {
	if tool.Parameters.ModelName == "" {
		tool.Parameters.ModelName = GetDefaultModel()
	}
	return tool
}
