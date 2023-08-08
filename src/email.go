package src

type EmailData struct {
	BrandName     string
	BrandUrl      string
	BrandImageUrl string
	BrandDomain   string

	Data any
}

func GenerateEmailData(data any, c *Config) EmailData {
	return EmailData{
		BrandName:     c.BrandName,
		BrandUrl:      c.SiteHost.String(),
		BrandImageUrl: "https://todo",
		BrandDomain:   c.SiteHost.Host,

		Data: data,
	}
}
