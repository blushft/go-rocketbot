package rc

type Client interface {
	Login() error
	Info() (*Info, error)
	GetRooms() (*Rooms, error)
	Me() (*Me, error)
	SendHook(string, Message) error
}
