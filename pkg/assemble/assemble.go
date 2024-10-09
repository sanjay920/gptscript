package assemble

import (
	"encoding/json"
	"io"

	"github.com/sanjay920/gptscript/pkg/types"
)

var Header = []byte("GPTSCRIPT!")

func Assemble(prg types.Program, output io.Writer) error {
	if _, err := output.Write(Header); err != nil {
		return err
	}
	return json.NewEncoder(output).Encode(prg)
}
