package cli

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/netograph/netograph-api/go/proto/ngapi"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func connect() (ngapi.NetographClient, context.Context, error) {
	tlsconf := &tls.Config{
		ServerName: "grpc.netograph.io",
	}
	addr := viper.GetString("addr")
	if strings.HasPrefix(addr, "localhost:") {
		tlsconf.InsecureSkipVerify = true
	}
	token := viper.GetString("token")
	if token == "" {
		return nil, nil, fmt.Errorf("no API token specified")
	}
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(credentials.NewTLS(tlsconf)),
	)
	if err != nil {
		return nil, nil, err
	}
	c := ngapi.NewNetographClient(conn)

	header := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	return c, ctx, nil
}
