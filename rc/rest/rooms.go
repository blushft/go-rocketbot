package rest

import (
	"encoding/json"

	"github.com/blushft/go-rocketbot/rc"
)

func (c *client) GetRooms() (*rc.Rooms, error) {
	resp, err := c.httpc.R().Get("rooms.get")
	if err != nil {
		return nil, err
	}

	r := new(rc.Rooms)

	if err = json.Unmarshal(resp.Body(), r); err != nil {
		return nil, err
	}
	return r, nil
}
