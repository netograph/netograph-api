package cli

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/spf13/viper"
	"github.com/netograph/netograph-api/go/proto/ngapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func connect() (ngapi.NetographClient, context.Context, error) {
	creds := credentials.NewTLS(&tls.Config{
		ServerName: "grpc.netograph.io",
	})
	addr := viper.GetString("addr")
	token := viper.GetString("token")
	if token == "" {
		return nil, nil, fmt.Errorf("no API token specified")
	}
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, nil, err
	}
	c := ngapi.NewNetographClient(conn)

	header := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	return c, ctx, nil
}
