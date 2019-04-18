package cli

import (
	"fmt"
	"os"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"golang.org/x/crypto/ssh/terminal"
)

func Output(cmd *cobra.Command, v proto.Message) error {
	cjson, err := cmd.Parent().PersistentFlags().GetBool("cjson")
	if err != nil {
		return err
	}
	color, err := cmd.Parent().PersistentFlags().GetBool("color")
	if err != nil {
		return err
	}
	term := terminal.IsTerminal(int(os.Stdout.Fd()))

	if term && !cjson {
		color = true
	}
	prettify := true
	if cjson {
		prettify = false
	}

	indent := "    "
	if cjson {
		indent = ""
	}
	m := jsonpb.Marshaler{
		EmitDefaults: false,
		Indent:       indent,
	}

	s, err := m.MarshalToString(v)
	if err != nil {
		return err
	}
	if prettify {
		s = string(pretty.Pretty([]byte(s)))
	}
	if color {
		s = string(pretty.Color([]byte(s), nil))
	}
	fmt.Println(s)
	return nil
}
