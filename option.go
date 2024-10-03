package osm

// Opts holds configuration options for requests.
type Opts struct {
	Locale string // Locale specifies the language for the response.
}

type optFunc func(*Opts)

// DefaultOpts provides the default options for requests.
var DefaultOpts = Opts{
	Locale: "en", // Default locale is English.
}

// WithLocale returns an optFunc that sets the locale option.
// locale specifies the language for the response.
func WithLocale(locale string) optFunc {
	return func(o *Opts) {
		o.Locale = locale
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
	return o
}
