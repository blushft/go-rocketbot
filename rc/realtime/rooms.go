package realtime

import (
	"github.com/Jeffail/gabs"
	"github.com/blushft/go-rocketbot/rc"
)

func (c *client) GetRooms() (*rc.Rooms, error) {
	resp, err := c.ddpc.Call("rooms/get", map[string]interface{}{"$date": 0})
	if err != nil {
		return nil, err
	}

	doc, err := gabs.Consume(resp.(map[string]interface{}))
	if err != nil {
		return nil, err
	}

	chans, err := doc.Children()
	up := make([]rc.Room, 0)

	for _, ch := range chans {
		up = append(up, rc.Room{
			ID:   stringOrZero(ch.Path("_id)").Data()),
			Name: stringOrZero(ch.Path("name").Data()),
			Type: stringOrZero(ch.Path("t").Data()),
		})
	}

	ret := &rc.Rooms{
		Update:  up,
		Remove:  nil,
		Success: true,
	}

	return ret, nil
}
