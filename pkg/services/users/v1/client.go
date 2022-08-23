package usersv1

import (
	"fmt"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	UsersClient
}

func NewClient(host string, port int, name string) (*Client, error) {
	c, err := grpc.Dial(fmt.Sprintf("dns:///%s:%d", host, port),
		grpc.WithInsecure(),
		grpc.WithUserAgent(name))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:        c,
		UsersClient: NewUsersClient(c),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
