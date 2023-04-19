package common

type LoaderOption struct {
	IDKey string
}

func getOptions(options []func(*LoaderOption)) *LoaderOption {
	// Init options
	opts := &LoaderOption{IDKey: "ID"}
	// Run options
	for _, o := range options {
		o(opts)
	}

	return opts
}

func WithIDKey(idKey string) func(*LoaderOption) {
	return func(lo *LoaderOption) { lo.IDKey = idKey }
}
