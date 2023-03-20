package lib

type CDNInterface interface {
	Purge(url string) error
	SoftPurge(url string) error
}

type CDN struct {
	APIToken string
}

func NewCDN(token string) CDNInterface {
	return &CDN{
		APIToken: token,
	}
}

func (c *CDN) Purge(url string) error {
	// TODO: https://docs.fastly.com/ja/guides/authenticating-api-purge-requests#api-%E3%83%88%E3%83%BC%E3%82%AF%E3%83%B3%E3%81%AB%E3%82%88%E3%82%8B-url-%E3%81%AE%E3%83%91%E3%83%BC%E3%82%B8
	return nil
}

func (c *CDN) SoftPurge(url string) error {
	// TODO: https://docs.fastly.com/ja/guides/soft-purges
	return nil
}
