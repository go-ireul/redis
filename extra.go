package redis

// Open combination of ParseURL and NewClient
func Open(url string) (c *Client, err error) {
	var f *Options
	if f, err = ParseURL(url); err == nil {
		c = NewClient(f)
	}
	return
}
