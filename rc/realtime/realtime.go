package realtime

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"

	"github.com/Jeffail/gabs"

	"github.com/blushft/go-rocketbot/rc"
	"github.com/gopackage/ddp"
)

type client struct {
	opts *rc.Options
	ddpc *ddp.Client
	auth *auth
}

func NewClient(opts ...rc.Option) rc.Client {
	dopts := rc.DefaultOptions()

	for _, opt := range opts {
		opt(dopts)
	}

	c := &client{
		opts: dopts,
	}

	c.initDDP()

	return c
}

func (c *client) initDDP() {
	u, _ := url.Parse(c.opts.URL)

	wsScheme := "ws"
	if u.Scheme == "https" {
		wsScheme = "wss"
	}

	wsURL := fmt.Sprintf("%s://%s/websocket", wsScheme, u.Host)

	ddpc := ddp.NewClient(wsURL, u.String())
	ddpc.SetSocketLogActive(c.opts.Debug)

	c.ddpc = ddpc

	c.Connect()
}

func (c *client) Connect() error {
	return c.ddpc.Connect()
}

func (c *client) Login() error {
	var req interface{}
	digest := sha256.Sum256([]byte(c.opts.Creds.Pass))
	req = ddpLoginRequest{
		User: ddpUser{Username: c.opts.Creds.User},
		Password: ddpPassword{
			Digest:    hex.EncodeToString(digest[:]),
			Algorithm: "sha-256",
		},
	}

	resp, err := c.ddpc.Call("login", req)
	if err != nil {
		return err
	}

	user := getUserFromData(resp.(map[string]interface{}))
	c.auth = &auth{
		id:    user.ID,
		token: user.Token,
	}

	return nil
}

func (c *client) Info() (*rc.Info, error) {
	panic("not implimented")
}

func (c *client) SendHook(ch string, msg rc.Message) error {
	panic("not implimented")
}

func (c *client) Reconnect() {
	c.ddpc.Reconnect()
}

func (c *client) Close() {
	c.ddpc.Close()
}

func (c *client) connectionAway() error {
	_, err := c.ddpc.Call("UserPresence:away")
	return err
}

func (c *client) connectionOnline() error {
	_, err := c.ddpc.Call("UserPresence:online")
	return err
}

type statusListner struct {
	listener func(int)
}

func (s statusListner) Status(status int) {
	s.listener(status)
}

func (c *client) AddStatusListener(listener func(int)) {
	c.ddpc.AddStatusListener(statusListner{listener: listener})
}

type auth struct {
	token string
	id    string
}

func randomID() string {
	return fmt.Sprintf("%f", rand.Float64())
}

func stringOrZero(i interface{}) string {
	if i == nil {
		return ""
	}

	switch i.(type) {
	case string:
		return i.(string)
	case float64:
		return fmt.Sprintf("%f", i.(float64))
	default:
		return ""
	}
}

type ddpLoginRequest struct {
	User     ddpUser     `json:"user"`
	Password ddpPassword `json:"password"`
}

type ddpTokenLoginRequest struct {
	Token string `json:"resume"`
}

type ddpUser struct {
	Username string `json:"username"`
}

type ddpPassword struct {
	Digest    string `json:"digest"`
	Algorithm string `json:"algorithm"`
}

func getUserFromData(data interface{}) *rc.User {
	document, _ := gabs.Consume(data)

	expires, _ := strconv.ParseFloat(stringOrZero(document.Path("tokenExpires.$date").Data()), 64)
	return &rc.User{
		ID:           stringOrZero(document.Path("id").Data()),
		Token:        stringOrZero(document.Path("token").Data()),
		TokenExpires: int64(expires),
	}
}
