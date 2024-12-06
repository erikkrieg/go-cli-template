// Greet is an overly complicated package that demonstrates some patterns, like
// like function options and table driven tests.
package greet

type GreetOptions struct {
	Subject *string
}

type Option func(*GreetOptions)

func WithSubject(subject string) Option {
	return func(o *GreetOptions) {
		o.Subject = &subject
	}
}

func Hello(opts ...Option) string {
	options := &GreetOptions{}
	for _, o := range opts {
		o(options)
	}
	if options.Subject == nil || *options.Subject == "" {
		subject := "World"
		options.Subject = &subject
	}

	return "Hello, " + *options.Subject
}
