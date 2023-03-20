package lib

type CDNInterface interface {
}

type CDN struct{}

func NewCDN() *CDN {
	return &CDN{}
}
