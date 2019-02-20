package rc

var (
	DefaultURL     = "http://localhost:3000"
	DefaultAPIPath = "/api/v1"
)

func DefaultOptions() *Options {
	return &Options{
		URL:      DefaultURL,
		APIPath:  DefaultAPIPath,
		Channels: make([]string, 0),
		Hooks:    make(map[string]string),
	}
}
