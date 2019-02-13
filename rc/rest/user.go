package rest

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"

	"github.com/blushft/go-rocketbot/rc"
)

func (c *client) Me() (*rc.Me, error) {
	resp, err := c.httpc.R().Get("me")
	if err != nil {
		return nil, err
	}

	r := new(rc.Me)

	if err = json.Unmarshal(resp.Body(), r); err != nil {
		return nil, err
	}
	spew.Dump(r)
	return r, nil
}
