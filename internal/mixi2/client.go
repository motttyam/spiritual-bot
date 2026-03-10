package mixi2

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/mixigroup/mixi2-application-sdk-go/auth"
	applicationapiv1 "github.com/mixigroup/mixi2-application-sdk-go/gen/go/social/mixi/application/service/application_api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Config struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
	APIAddress   string
}

type Client struct {
	authenticator *auth.Authenticator
	apiClient     applicationapiv1.ApplicationServiceClient
	conn          *grpc.ClientConn
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.ClientID == "" {
		return nil, fmt.Errorf("MIXI2_CLIENT_ID is required")
	}
	if cfg.ClientSecret == "" {
		return nil, fmt.Errorf("MIXI2_CLIENT_SECRET is required")
	}
	if cfg.TokenURL == "" {
		return nil, fmt.Errorf("MIXI2_TOKEN_URL is required")
	}
	if cfg.APIAddress == "" {
		return nil, fmt.Errorf("MIXI2_API_ADDRESS is required")
	}

	authenticator, err := auth.NewAuthenticator(cfg.ClientID, cfg.ClientSecret, cfg.TokenURL)
	if err != nil {
		return nil, fmt.Errorf("create authenticator: %w", err)
	}

	conn, err := grpc.NewClient(
		cfg.APIAddress,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
	)
	if err != nil {
		return nil, fmt.Errorf("create grpc connection: %w", err)
	}

	return &Client{
		authenticator: authenticator,
		apiClient:     applicationapiv1.NewApplicationServiceClient(conn),
		conn:          conn,
	}, nil
}

func (c *Client) Post(ctx context.Context, text string) error {
	if text == "" {
		return fmt.Errorf("post text is empty")
	}
	if len([]rune(text)) > 149 {
		return fmt.Errorf("post text too long: %d", len([]rune(text)))
	}

	authCtx, err := c.authenticator.AuthorizedContext(ctx)
	if err != nil {
		return fmt.Errorf("authorize context: %w", err)
	}

	_, err = c.apiClient.CreatePost(authCtx, &applicationapiv1.CreatePostRequest{Text: text})
	if err != nil {
		return fmt.Errorf("create post: %w", err)
	}
	return nil
}

func (c *Client) Close() error {
	if c.conn == nil {
		return nil
	}
	return c.conn.Close()
}
