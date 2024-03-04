package lib

import "github.com/fastly/go-fastly/v7/fastly"

type CDNInterface interface {
	Purge(url string) error
	SoftPurge(url string) error
}

type CDN struct {
	Client *fastly.Client
}

func NewCDN(token string) (CDNInterface, error) {
	client, err := fastly.NewClient(token)
	if err != nil {
		return nil, err
	}
	return &CDN{
		Client: client,
	}, nil
}

func (c *CDN) Purge(url string) error {
	purgeInput := &fastly.PurgeInput{
		URL:  url,
		Soft: false,
	}
	_, err := c.Client.Purge(purgeInput)
	if err != nil {
		return err
	}
	return nil
}

func (c *CDN) SoftPurge(url string) error {
	purgeInput := &fastly.PurgeInput{
		URL:  url,
		Soft: true,
	}
	_, err := c.Client.Purge(purgeInput)
	if err != nil {
		return err
	}
	return nil
}
