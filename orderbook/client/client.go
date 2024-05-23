package client

type Client struct {
	id string
}

type IClient interface {
	GetID() string
}

func NewClient(id string) *Client {
	return &Client{id: id}
}

func (c *Client) GetID() string {
	return c.id
}
