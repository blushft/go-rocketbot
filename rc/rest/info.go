package rest

import (
	"encoding/json"

	"github.com/blushft/go-rocketbot/rc"
)

func (c *client) Info() (*rc.Info, error) {
	resp, err := c.httpc.R().Get("info")
	if err != nil {
		return nil, err
	}

	r := new(rc.Info)

	if err = json.Unmarshal(resp.Body(), r); err != nil {
		return nil, err
	}
	return r, nil
}
