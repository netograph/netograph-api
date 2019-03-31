package cli

import (
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
)

type Outputter interface {
	Summary() string
}

func Output(cmd *cobra.Command, v proto.Message) error {
	cjson, err := cmd.Parent().PersistentFlags().GetBool("cjson")
	if err != nil {
		return err
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
	fmt.Println(s)
	return nil
}
