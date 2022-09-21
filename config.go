package wechatqy

func (c *Client) Config(key string) *Client {
	c.config.key = key
	return c
}
