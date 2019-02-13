package rest

import (
	"testing"
	"time"

	"github.com/blushft/go-rocketbot/rc"
)

func TestClient_Login(t *testing.T) {
	tests := []struct {
		name    string
		c       rc.Client
		h       rc.Message
		wantErr bool
	}{
		{
			name: "client login test",
			c: NewClient(
				rc.URL("http://192.168.99.100:3000"),
				rc.Credentials("botty", "123456", "ldap"),
				rc.RegisterHook("testhook", "XgtW2q68dJrp5t2wL/Nq7mWYSKnF3DDqbpWWuwTSrf9Jt8GE84griuKX4e4kzhzHpT"),
			),
			h: rc.Message{
				Text:    "this is a test https://via.placeholder.com/25",
				Channel: "#talk,#general",
				Attachments: []rc.Attachment{
					{
						Color:     "blue",
						Timestamp: time.Now(),
						ImageURL:  "https://via.placeholder.com/150",
						Title:     "The title",
						Text:      "The text",
						/* Fields: []rc.Field{
							{
								Title: "Field 1",
								Value: "Value",
								Short: true,
							},
							{
								Title: "Field 2",
								Value: "Value",
								Short: true,
							},
						}, */
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Login(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
			if i, err := tt.c.Info(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Info() got = %v, err = %v, wanterr = %v", i, err, tt.wantErr)
			}
			if r, err := tt.c.GetRooms(); (err != nil) != tt.wantErr {
				t.Errorf("Client.GetRooms() got = %v, err = %v, wanterr = %v", r, err, tt.wantErr)
			}
			if me, err := tt.c.Me(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Me() got = %v, err = %v, wanterr = %v", me, err, tt.wantErr)
			}
			if err := tt.c.SendHook("testhook", tt.h); (err != nil) != tt.wantErr {
				t.Errorf("Client.SendHook() err = %v, wanterr = %v", err, tt.wantErr)
			}
		})
	}
}
