package config

import "fmt"

type Config struct {
	Url            string
	Authentication bool
	Token          string
}

func NewConfig(url, token string, auth bool) Config {
	return Config{
		Url:            url,
		Authentication: auth,
		Token:          token,
	}
}

func (c *Config) String() string {
	return fmt.Sprintf("URL: %s, Auth: %v", c.Url, c.Authentication)
}
