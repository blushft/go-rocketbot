package rest

import (
	"encoding/json"
	"errors"
	"net/url"
	"path"

	"github.com/blushft/go-rocketbot/rc"
	"github.com/blushft/go-rocketbot/rc/webhook"
	resty "gopkg.in/resty.v1"
)

var (
	defaultURL     = "http://localhost:3000"
	defaultAPIPath = "/api/v1"
)

type client struct {
	opts  rc.Options
	auth  *auth
	httpc *resty.Client
	hookc *webhook.Client
}

func NewClient(opts ...rc.Option) rc.Client {
	dopts := rc.Options{
		URL:      defaultURL,
		APIPath:  defaultAPIPath,
		Channels: make([]string, 0),
		Hooks:    make(map[string]string),
	}

	for _, opt := range opts {
		opt(&dopts)
	}

	c := &client{
		opts:  dopts,
		httpc: resty.New(),
		hookc: webhook.NewClient(opts...),
	}

	c.initHttpC()

	return c
}

func (c *client) initHttpC() {
	u, _ := url.Parse(c.opts.URL)

	u.Path = path.Join(u.Path, c.opts.APIPath)
	c.httpc.HostURL = u.String()
}

type auth struct {
	token string
	id    string
}

func (c *client) Login() error {
	if c.opts.Creds.User == "" {
		return errors.New("no credentials")
	}

	resp, err := c.httpc.R().
		SetFormData(map[string]string{
			"user":     c.opts.Creds.User,
			"password": c.opts.Creds.Pass,
		}).
		Post("login")
	if err != nil {
		return err
	}

	r := new(logonResponse)
	if err := json.Unmarshal(resp.Body(), r); err != nil {
		return err
	}

	c.auth = &auth{
		id:    r.Data.UserID,
		token: r.Data.Token,
	}

	c.setAuthHeader()
	return nil
}

func (c *client) SendHook(hook string, msg rc.Message) error {
	return c.hookc.SendHook(hook, msg)
}

func (c *client) setAuthHeader() {
	c.httpc.SetHeaders(map[string]string{
		"X-Auth-Token": c.auth.token,
		"X-User-Id":    c.auth.id,
	})
}

type logonResponse struct {
	Status string    `json:"status,omitempty"`
	Data   logonData `json:"data,omitempty"`
}

type logonData struct {
	Token  string `json:"authToken,omitempty"`
	UserID string `json:"userID,omitempty"`
}

type Pagination struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type Status struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`

	Status  string `json:"status"`
	Message string `json:"message"`
}
