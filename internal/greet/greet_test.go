package greet

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		args []Option
		want string
	}{
		{
			name: "No options",
			args: []Option{},
			want: "Hello, World",
		},
		{
			name: "Some options",
			args: []Option{
				WithSubject("Moon"),
			},
			want: "Hello, Moon",
		},
		{
			name: "Colliding options",
			args: []Option{
				WithSubject("Moon"),
				WithSubject("Spoon"),
			},
			want: "Hello, Spoon",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.args) > 0 {
				got = Hello(tt.args...)
			} else {
				got = Hello()
			}
			if tt.want != got {
				t.Errorf("Wanted %s but got %s\n", tt.want, got)
			}
		})
	}
}
