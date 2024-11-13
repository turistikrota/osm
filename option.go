package osm

// Opts holds configuration options for requests.
type Opts struct {
	Locale    string // Locale specifies the language for the response.
	UserAgent string // UserAgent specifies the User-Agent header for the request.
}

type optFunc func(*Opts)

// DefaultOpts provides the default options for requests.
var DefaultOpts = Opts{
	Locale:    "en",     // Default locale is English.
	UserAgent: "Chrome", // Default User-Agent is Chrome.
}

// WithLocale returns an optFunc that sets the locale option.
// locale specifies the language for the response.
func WithLocale(locale string) optFunc {
	return func(o *Opts) {
		o.Locale = locale
	}
}

// WithUserAgent returns an optFunc that sets the User-Agent option.
// userAgent specifies the User-Agent header for the request.
func WithUserAgent(userAgent string) optFunc {
	return func(o *Opts) {
		o.UserAgent = userAgent
	}
}

func mergeOpts(optsList []optFunc) *Opts {
	o := &Opts{}
	for _, opt := range optsList {
		opt(o)
	}
	if o.Locale == "" {
		o.Locale = DefaultOpts.Locale
	}
	if o.UserAgent == "" {
		o.UserAgent = DefaultOpts.UserAgent
	}
	return o
}
