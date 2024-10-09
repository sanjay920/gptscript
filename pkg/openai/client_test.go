package openai

import (
	"testing"

	"github.com/hexops/autogold/v2"
	"github.com/hexops/valast"
	openai "github.com/sanjay920/chat-completion-client"
	"github.com/sanjay920/gptscript/pkg/types"
)

func Test_appendMessage(t *testing.T) {
	autogold.Expect(types.CompletionMessage{Content: []types.ContentPart{
		{ToolCall: &types.CompletionToolCall{
			Index: valast.Ptr(0),
			Function: types.CompletionFunctionCall{
				Name:      "foo",
				Arguments: "bar",
			},
		}},
		{ToolCall: &types.CompletionToolCall{
			Index: valast.Ptr(1),
			Function: types.CompletionFunctionCall{
				Name:      "foo",
				Arguments: "bar",
			},
		}},
	}}).Equal(t, appendMessage(types.CompletionMessage{}, openai.ChatCompletionStreamResponse{
		Choices: []openai.ChatCompletionStreamChoice{
			{
				Delta: openai.ChatCompletionStreamChoiceDelta{
					ToolCalls: []openai.ToolCall{
						{
							Function: openai.FunctionCall{
								Name:      "foo",
								Arguments: "bar",
							},
						},
						{
							Function: openai.FunctionCall{
								Name:      "foo",
								Arguments: "bar",
							},
						},
					},
				},
			},
		},
	}))
}
