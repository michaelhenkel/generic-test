package client

import (
	"fmt"

	"github.com/michaelhenkel/generic-test/models"
)

type Client struct{}

func (c *Client) Add(resource models.ResourceA) {
	fmt.Printf("adding resource %+v", resource)
}

type ClientInterface interface {
	Add()
}
