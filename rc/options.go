package rc

type Options struct {
	URL      string
	APIPath  string
	HookPath string
	Channels []string
	Hooks    map[string]string

	Debug bool
	Creds Creds
}

type Option func(*Options)

type Creds struct {
	User   string
	Pass   string
	Token  string
	Method string
}

func URL(url string) Option {
	return func(o *Options) {
		o.URL = url
	}
}

func APIPath(path string) Option {
	return func(o *Options) {
		o.APIPath = path
	}
}

func HookPath(path string) Option {
	return func(o *Options) {
		o.HookPath = path
	}
}

func Subscribe(channel string) Option {
	return func(o *Options) {
		o.Channels = append(o.Channels, channel)
	}
}

func RegisterHook(name, hook string) Option {
	return func(o *Options) {
		o.Hooks[name] = hook
	}
}

func Debug(d bool) Option {
	return func(o *Options) {
		o.Debug = d
	}
}

func Credentials(user, pass, method string) Option {
	return func(o *Options) {
		o.Creds = Creds{
			User:   user,
			Pass:   pass,
			Method: method,
		}
	}
}
