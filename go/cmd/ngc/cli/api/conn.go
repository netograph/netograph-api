package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/netograph/netograph-api/go/proto/ngapi/userapi"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func connect() (*grpc.ClientConn, context.Context, error) {
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

	header := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	return conn, ctx, nil
}

func connectUser() (userapi.UserClient, context.Context, error) {
	conn, ctx, err := connect()
	if err != nil {
		return nil, nil, err
	}
	return userapi.NewUserClient(conn), ctx, nil
}

func connectDset() (dsetapi.DsetClient, context.Context, error) {
	conn, ctx, err := connect()
	if err != nil {
		return nil, nil, err
	}
	return dsetapi.NewDsetClient(conn), ctx, nil
}
