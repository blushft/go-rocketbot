package webhook

import (
	"errors"
	"net/url"
	"path"

	"github.com/blushft/go-rocketbot/rc"

	resty "gopkg.in/resty.v1"
)

var (
	defaultURL      = "http://localhost:3000"
	defaultHookPath = "/hooks"
)

type Client struct {
	opts  *rc.Options
	httpc *resty.Client
}

func NewClient(opts ...rc.Option) *Client {
	dopts := rc.DefaultOptions()

	for _, opt := range opts {
		opt(dopts)
	}

	c := &Client{
		opts:  dopts,
		httpc: resty.New(),
	}

	c.initHttpC()

	return c
}

func (c *Client) initHttpC() {
	u, _ := url.Parse(c.opts.URL)

	u.Path = path.Join(u.Path, c.opts.HookPath)
	c.httpc.HostURL = u.String()
}

func (c *Client) SendHook(hook string, msg rc.Message) error {
	p, ok := c.opts.Hooks[hook]
	if !ok {
		return errors.New("hook not found")
	}
	_, err := c.httpc.R().
		SetBody(msg).
		Post(p)
	if err != nil {
		return err
	}
	return nil
}
