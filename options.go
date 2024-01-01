package opencage

import "time"

type Option func(*Client)

func WithEndpoint(endpoint string) Option {
	return func(client *Client) {
		client.endpoint = endpoint
	}
}

func WithGlobalTimeout(duration time.Duration) Option {
	return func(client *Client) {
		client.httpClient.Timeout = duration
	}
}
