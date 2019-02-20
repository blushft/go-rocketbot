package realtime

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/blushft/go-rocketbot/rc"
)

var (
	TRTClient = NewClient(
		rc.URL("http://192.168.99.100:3000"),
		rc.Credentials("testuser", "123456", "email"),
		rc.Debug(true),
	)
)

func TestGetRooms(t *testing.T) {

	tests := []struct {
		name string
		want *rc.Rooms
	}{
		{
			name: "test getrooms",
			want: &rc.Rooms{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TRTClient.Login(); err != nil {
				t.Errorf("error Login(): %v", err)
			}
			rooms, err := TRTClient.GetRooms()
			if err != nil {
				t.Errorf("error GetRooms(): %v", err)
			}
			spew.Dump(rooms)
		})
	}
}
