package linkding

type LinkdingClientOptions struct {
	Token string
	Host  string
}

type LinkdingClient struct {
	token string
	host  string
}

func NewLinkdingClient(opts LinkdingClientOptions) *LinkdingClient {
	client := &LinkdingClient{
		token: opts.Token,
		host:  opts.Host,
	}

	return client
}
