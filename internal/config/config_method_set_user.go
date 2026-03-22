package config

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}
