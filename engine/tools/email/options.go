package email

type ClientOptions func(*Client)

func WithHost(hostname string) ClientOptions {
	return func(c *Client) {
		c.server.Host = hostname
	}
}

func WithPassword(password string) ClientOptions {
	return func(c *Client) {
		c.server.Password = password
	}
}

func WithUsername(username string) ClientOptions {
	return func(c *Client) {
		c.server.Username = username
	}
}

func WithPort(port int) ClientOptions {
	return func(c *Client) {
		c.server.Port = port
	}
}

func WithEncryption(encryption string) ClientOptions {
	return func(c *Client) {
	}
}

func WithSenderEmail(senderEmail string) ClientOptions {
	return func(c *Client) {
		c.email.SetFrom(senderEmail)
	}
}
